CGO_CFLAGS="-I/Users/zhs007/go/src/svn.heyalgo.io/slots/rocksdb/include" \
CGO_LDFLAGS="-L/Users/zhs007/go/src/svn.heyalgo.io/slots/rocksdb -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \
  go build ./