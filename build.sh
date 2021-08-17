
echo "Building fabric-agent..."
pushd agent
if [ -f "fabric-agent" ]
then
    echo "fabric-agent already exists, no need to rebuild."
else
    ./build.sh
fi
popd

echo "Building fabric-server..."
pushd server
if [ -f "fabric-server" ]
then
    echo "fabric-server already exists, no need to rebuild."
else
    ./build.sh
fi    
popd

echo "Building fabric-explorer..."
pushd explorer
if [ -f "fabric-explorer" ]
then
    echo "fabric-explorer already exists, no need to rebuild."
else
    ./build.sh
fi
popd