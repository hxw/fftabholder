# fftabholder

This is a simple http server for use with
Tree Style Tab Firefox extension


# Installation on FreeBSD

Install and set up Golang aonfigure its environment variables

~~~~~
go get github.com/hxw/fftabholder
cd "${GOPATH}/src/github.com/hxw/fftabholder"
make
sudo make install
sudo sysrc fftabholder_enable=YES
sudo service fftabholder start
~~~~~

# Using with Firefox

Open a new tab and prowse to:

~~~~~
http://127.0.0.1:8016/Text%20Documents
~~~~~

A tab called "Text Documents" will be created, simply drag and drop
related tabs ad children of this top level tab.

The form on the placeholder tab will open a child tab using the url entered,
however if the first character is a period(.) then it will open as a
child placeholder tab.

# To remove

~~~~~
service fftabholder stop
sudo make deinstall
sudo sysrc -x fftabholder_enable
~~~~~
