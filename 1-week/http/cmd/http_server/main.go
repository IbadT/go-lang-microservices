package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-chi/chi"
)

const (
	baseUrl       = "localhost:8081"
	createPostfix = "/notes"
	getPostfix    = "/notes/%d"
)

// используем `json:""` для того, чтобы легче было работать с json.Marshal и Unmarshal
type NoteInfo struct {
	Title    string `json:"title"`
	Context  string `json:"context"`
	Author   string `json:"author"`
	IsPublic bool   `json:"is_public"`
}

type Note struct {
	ID        int64     `json:"id"`
	Info      NoteInfo  `json:"info"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	r := chi.NewRouter()

	r.Post(createPostfix, createNoteHandler)
	r.Get(getPostfix, getNoteHandler)

	// блокирующий вызов ListenAndServe
	err := http.ListenAndServe(baseUrl, r)
	if err != nil {
		log.Fatal(err)
	}
}

func parseNoteID(idStr string) (int64, error) {
	// 10 - десятичная система
	// 64 - int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

type SyncMap struct {
	elems map[int64]*Note
	// мы допускаем, что будет конкурентный доступ к ресурсу
	// тк эти endpointы могут вызывать много клиентов
	// поэтому делаем условие, что будет большая нагрузка на чтение
	m sync.RWMutex
}

// глобальный объект
var notes = &SyncMap{
	elems: make(map[int64]*Note),
}

// ResponseWriter - объект, в который мы будем записывать наш ответ
// Request - те данные, которые придут из сервера, будет хранить request
func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	// создаем default объект
	info := &NoteInfo{}
	// NewDecoder - декодируем в нашу структуру из байтов
	// Decode -
	if err := json.NewDecoder(r.Body).Decode(info); err != nil {
		http.Error(w, "Fatal to decode note data", http.StatusBadRequest)
		return
	}

	rand.Seed(time.Now().UnixNano())
	now := time.Now()

	note := &Note{
		ID:        rand.Int63(),
		Info:      *info,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// настраиваем заголовки
	w.Header().Set("Content-Type", "application/json")
	// настраиваем статус ответа
	w.WriteHeader(http.StatusCreated)
	// инкодим нашу структуру в json и наши байты складываются в наш w - ResponseWriter
	// и клиенту прилетит весь объект целиком note
	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "Failed to encode note data", http.StatusInternalServerError)
		return
	}

	// блокируем
	notes.m.Lock()

	//
	defer notes.m.Unlock()

	notes.elems[note.ID] = note
}

func getNoteHandler(w http.ResponseWriter, r *http.Request) {
	// получаем из request свойство с id
	// передаем объект request-а и название нашего параметра "id"
	noteID := chi.URLParam(r, "id")
	id, err := parseNoteID(noteID)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	// RLock - запрещает в момент, когда этот lock, что-то записывать
	// запрещена запись например(если взят элемент с id = 5, перезаписать этот id)
	// и чтобы получить этот объект
	notes.m.RLock()
	defer notes.m.RUnlock()

	note, ok := notes.elems[id]
	if !ok {
		http.Error(w, "Note id not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "Failed to encode note data", http.StatusInternalServerError)
		return
	}

}

// defer и os.Exit(1)
