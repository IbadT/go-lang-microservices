# gRPC
в папку api складывают описание внешних api

1 изначально нужно описать все в api/note_v1/note.proto

2 потом написать Makefile

3 make install-deps

4 make get-deps

5 make generate
в этом описании можно добавить все остальные grpc файлы(generate-note-api)

после этого будет создан 