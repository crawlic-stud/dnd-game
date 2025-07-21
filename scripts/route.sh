#!bin/bash

name=$1
test_name=${name}_test
camelName=$(echo "$name" | sed -r 's/(.)_+(.)/\1\U\2/g;s/^[a-z]/\U&/')

echo "" > internal/server/router/$name.go

echo "package router" >> internal/server/router/$name.go
echo "" >> internal/server/router/$name.go

echo "import (" >> internal/server/router/$name.go
echo "	\"net/http\"" >> internal/server/router/$name.go
echo ")" >> internal/server/router/$name.go
echo "" >> internal/server/router/$name.go

echo "func (api *Router) $camelName(w http.ResponseWriter, r *http.Request) {" >> internal/server/router/$name.go
echo "" >> internal/server/router/$name.go
echo "}" >> internal/server/router/$name.go

echo "" > internal/server/router/$test_name.go

echo "package router" >> internal/server/router/$test_name.go
echo "" >> internal/server/router/$test_name.go

echo "import (" >> internal/server/router/$test_name.go
echo "	\"testing\"" >> internal/server/router/$test_name.go
echo ")" >> internal/server/router/$test_name.go
echo "" >> internal/server/router/$test_name.go

echo "func  Test$camelName(t *testing.T) {" >> internal/server/router/$test_name.go
echo "// TODO" >> internal/server/router/$test_name.go
echo "}" >> internal/server/router/$test_name.go