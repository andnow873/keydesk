#!/bin/sh

### Create brigades

# * Check if brigade already exists
# * Create system user
# * Create homedir

# * Create json datafile
# * Create special brigadier wg-user

# * Activate keydesk systemD units
# * Activate stats systemD units

# * Send brigadier config

# creating brigade and brigadier app.

BASE_HOME_DIR="/home"
BASE_STATS_DIR="/var/lib/vgstats"
BRIGADE_MAKER_APP_PATH="/opt/vgkeydesk/createbrigade"
KEYDESK_APP_PATH="/opt/vgkeydesk/keydesk"

VGCERT_GROUP="vgcert"
VGSTATS_GROUP="vgstats"

spinlock="`[ ! -z \"${TMPDIR}\" ] && echo -n \"${TMPDIR}/\" || echo -n \"/tmp/\" ; echo \"vgbrigade.spinlock\"`"
trap "rm -f \"${spinlock}\" 2>/dev/null" EXIT
while [ -f "${spinlock}" ] ; do
    sleep 0.1
done
touch "${spinlock}" 2>/dev/null

set -e

printdef () {
        echo "Usage: $0 <brigabe_id_encoded> <endpoint IPv4> <CGNAT IPv4> <IPv6 ULA> <DNS IPv4> <DNS IPv6> <keydesk IPv6> <B1rigadier Name :: base64> <Person Name :: base64> <Person Desc :: base64> <Person URL :: base64>" >&2
        exit 1
}

if [ -z "${1}" ] || \
         [ -z "${2}" ] || \
         [ -z "${3}" ] || \
         [ -z "${4}" ] || \
         [ -z "${5}" ] || \
         [ -z "${6}" ] || \
         [ -z "${7}" ] || \
         [ -z "${8}" ] || \
         [ -z "${9}" ] || \
         [ -z "${10}" ] || \
         [ -z "${11}" ]; then 
        printdef
fi

brigade_id=${1}
endpoint_ip4=${2}
ip4_cgnat=${3}
ip6_ula=${4}
dns_ip4=${5}
dns_ip6=${6}
keydesk_ip6=${7}
brigadier_name=${8}
person_name=${9}
person_desc=${10}
person_url=${11}

shift 11

chunked=""
port="0"
domain=""

for i in "$@";
do
    case $i in
        [0-9]*)
                if [ "$i" -ge 1024 ] && [ "$i" -le 65535 ]; then
                        port="$i"
                fi
                ;;
        *.*)
                if printf "%s" "$i" | grep -E '^([a-z0-9_]+(-[a-z0-9_]+)*\.)+[a-z0-9_]+([a-z0-9_-]+)$' > /dev/null; then
                        domain="$i"
                fi
        ;;
        *)
                if [ "$i" = "chunked" ]; then
                        chunked="-ch"
                fi
        ;;
    esac
done

# * Check if brigade is exists

if [ -s "${BASE_HOME_DIR}/${brigade_id}/created" ]; then
        echo "Brigade ${brigade_id} already exists" >&2
        exit 1
fi

# * Create system user

useradd -p '*' -G "${VGCERT_GROUP}" -M -s /usr/sbin/nologin -d "${BASE_HOME_DIR}/${brigade_id}" "${brigade_id}"
install -o "${brigade_id}" -g "${brigade_id}" -m 0700 -d "${BASE_HOME_DIR}/${brigade_id}"
install -o "${brigade_id}" -g "${VGSTATS_GROUP}" -m 710 -d "${BASE_STATS_DIR}/${brigade_id}"

# Create json datafile

if ! out=$(sudo -u "${brigade_id}" -g "${brigade_id}" "${BRIGADE_MAKER_APP_PATH}" \
        -ep4 "${endpoint_ip4}" \
        -dns4 "${dns_ip4}" \
        -dns6 "${dns_ip6}" \
        -int4 "${ip4_cgnat}" \
        -int6 "${ip6_ula}" \
        -kd6 "${keydesk_ip6}" \
        -p "$port" \
        -dn "$domain" \
        ); then
        echo "Can't create brigade: ${out}" >&2

        exit 1
fi

if ! wgconf=$(sudo -u "${brigade_id}" -g "${brigade_id}" "${KEYDESK_APP_PATH}" \
        "${chunked}" \
        -name "${brigadier_name}" \
        -person "${person_name}" \
        -desc "${person_desc}" \
        -url "${person_url}"); then
        echo "Can't create brigadier: ${wgconf}" >&2

        exit 1
fi

# * Activate keydesk systemD units

systemd_vgkeydesk_instance="vgkeydesk@${brigade_id}"

# create dir for custom config
# https://www.freedesktop.org/software/systemd/man/systemd.unit.html
systemd_vgkeydesk_conf_dir="/etc/systemd/system/${systemd_vgkeydesk_instance}.socket.d"

#shellcheck disable=SC2174
mkdir -p "${systemd_vgkeydesk_conf_dir}" -m 0755

# it;s necessary to listen certain IP

# calculated listen IPv6 
listen_ip6=$(echo "${endpoint_ip4}" | sed 's/\./\n/g' | xargs printf 'fdcc:%02x%02x:%02x%02x::2' | sed 's/:0000/:/g' | sed 's/:00/:/g')

cat << EOF > "${systemd_vgkeydesk_conf_dir}/listen.conf"
[Socket]
ListenStream = [${listen_ip6}]:80
ListenStream = [${listen_ip6}]:443
EOF

systemctl -q enable "${systemd_vgkeydesk_instance}.socket" "${systemd_vgkeydesk_instance}.service"

# Start systemD services
systemctl -q start "${systemd_vgkeydesk_instance}.socket" "${systemd_vgkeydesk_instance}.service"

# * Activate stats systemD units

systemd_vgstats_instance="vgstats@${brigade_id}"
systemctl -q enable "${systemd_vgstats_instance}.service"
systemctl -q start "${systemd_vgstats_instance}.service"

# Print brigadier config
echo "${wgconf}"

date -u +"%Y-%m-%dT%H:%M:%S" > "${BASE_HOME_DIR}/${brigade_id}/created"
