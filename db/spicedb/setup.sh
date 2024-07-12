brew install authzed/tap/spicedb authzed/tap/zed

spicedb serve --grpc-preshared-key dasfaseraefasdfasdf

zed context set gitclone grpc.authzed.com:443 dasfaseraefasdfasdf

zed schema write schema.zed

zed relationship create izumarth/repo:1 writer izumarth/user:izumarth
zed relationship create izumarth/repo:1 reader izumarth/user:azuma

zed permission check izumarth/repo:1 read izumarth/user:izumarth # true
zed permission check izumarth/repo:1 write izumarth/user:izumarth # true
zed permission check izumarth/repo:1 read  izumarth/user:azuma  # true
zed permission check izumarth/repo:1 write izumarth/user:azuma  # fals