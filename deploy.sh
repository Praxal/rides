#!/bin/bash
sshcmd="ssh -t praxal@app.p4family.com"
$sshcmd screen -S "deployment" /home/praxal/rides/prod_deploy.sh
