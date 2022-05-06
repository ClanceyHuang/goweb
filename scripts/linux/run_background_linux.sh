# Executable file authorization, and run in the background.
cd ../../cmd

if [ -x GoWeb-linux ]
  then
    echo "start server by script"
  else
    chmod a+x GoWeb-linux
fi

nohup ./GoWeb-linux > server.out &


