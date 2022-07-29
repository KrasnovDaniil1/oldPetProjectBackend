package variable

const CREATE_FOLDER = "images"

const OPEN_POSTGRES = "user=user dbname=db password=pass sslmode=disable"

const CREATE_TABLE = "create table if not exists images (id SERIAL, tags varchar(10)[], filename text);"

const UPDATE_TABLE = "insert into images(tags, filename) values($1, $2)"

const GET_ALL_CARD = "SELECT * FROM images LIMIT $1 OFFSET $2"

const DELETE_CARD_BY_ID = "DELETE from images where id = $1"

const GET_LAST_ID = "SELECT id FROM images ORDER BY id DESC LIMIT 1"

const GET_SEARCH_CARD = "SELECT * FROM images where $1 = any(tags) LIMIT $2 OFFSET $3;"

const GET_ALL_TAGS = "SELECT id, tags FROM images"

const GET_SEARCH_TAG = "SELECT * FROM images where $1 = any(tags) LIMIT 1"

var VALID_IMAGE_FORMAT = []string{".png", ".jpg", ".jpeg"}
