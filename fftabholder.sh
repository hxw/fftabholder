#!/bin/sh
#
# $FreeBSD$
#
# PROVIDE: fftabholder
# REQUIRE: LOGIN cleanvar
# KEYWORD: shutdown
#
# Add the following lines to /etc/rc.conf.local or /etc/rc.conf
# to enable this service:
#
# fftabholder_enable (bool):	Set to NO by default.
#			Set it to YES to enable fftabholder.

. /etc/rc.subr

# check_process()
# {
#         /bin/pgrep -f 'fftabholder'
# }

fftabholder_prestart()
{
        # fftabholder_flags gets applied too early if we don't do this.
        # I didn't want to force people to update their rc.conf files
        # and change the fftabholder_flags to something else.
        rc_flags=""
}


name=fftabholder
rcvar=fftabholder_enable

load_rc_config $name

: ${fftabholder_enable="NO"}

start_precmd="fftabholder_prestart"
pidfile=/var/run/fftabholder.pid
procname=/usr/local/sbin/fftabholder
command=/usr/sbin/daemon
command_args=" -f -u www -p ${pidfile} ${procname} ${fftabholder_flags}"
#required_files=/usr/local/etc/fftabholder.conf

run_rc_command "$1"
