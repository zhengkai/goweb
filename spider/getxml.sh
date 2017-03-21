#!/bin/bash
cd `dirname $0`

# wget http://wot.kongzhong.com/wiki/xml/proList.xml
wget -q 'https://api.worldoftanks.ru/wot/encyclopedia/vehicles/?application_id=demo&language=zh-cn' -O vehicles.json
