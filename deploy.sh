#!/bin/bash
sshcmd="ssh -t praxal@app.p4family.com"
$sshcmd screen -S "deployment" /home/praxal/app/prod_deploy.sh
