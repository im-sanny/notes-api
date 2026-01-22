package store

import (
	"notes/model"
	"time"
)

var (
	Data = []model.Note{
		{
			ID:        1,
			Title:     "Grocery List - Weekend Plan",
			Content:   "Buy:\n- Milk (2L)\n- Eggs (dozen)\n- Bananas (bunch)\n- Chicken breast 1kg\n- Tomatoes, onions, garlic\n- Bread\n- Coffee beans\n\nAlso need:\n• Toilet paper\n• Dish soap\n• Cat food (salmon flavor)",
			CreatedAt: time.Date(2026, time.January, 22, 8, 5, 0, 0, time.UTC),
			UpdatedAt: time.Date(2026, time.January, 22, 9, 40, 30, 0, time.UTC),
		},
		{
			ID:        2,
			Title:     "Project Ideas for Q1 2026",
			Content:   "1. Personal finance dashboard (Go + HTMX + SQLite)\n2. Recipe recommender based on fridge ingredients\n3. Minimalist habit tracker with streak visualization\n4. Learn WebAssembly + Go (tiny game or image filter?)\n5. Rewrite old Python script → Go CLI tool\n\nPriorities:\n- Finance dashboard first (most daily value)\n- Start small, ship fast\n\nDeadline ideas: end of Feb for MVP of #1",
			CreatedAt: time.Date(2026, time.January, 22, 8, 5, 0, 0, time.UTC),
			UpdatedAt: time.Date(2026, time.January, 22, 9, 40, 30, 0, time.UTC),
		},
	}

	NextId int64 = 0
)
