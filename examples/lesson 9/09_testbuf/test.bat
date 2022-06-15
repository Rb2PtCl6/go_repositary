
@echo off

cd test_buf
go build test_buf.go
move test_buf.exe ..\ >null
cd ..

cd test_buff
go build test_buff.go
move test_buff.exe ..\ >null
cd ..

cd test_nobuf
go build test_nobuf.go
move test_nobuf.exe ..\ >null
cd ..

cd test_nobuff
go build test_nobuff.go
move test_nobuff.exe ..\ >null
cd ..

echo test_buf
time <%0
test_buf.exe
echo test_buff
time <%0
test_buff.exe
echo test_nobuf
time <%0
test_nobuf.exe
echo test_nobuff
time <%0
test_nobuff.exe
time <%0

del test*.exe
del num.res
pause

