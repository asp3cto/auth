package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate /home/aspecto/go/bin/minimock -i UserRepository -o ./mocks/ -s "_minimock.go"
