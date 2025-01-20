package service

import "strings"

func FilterProductURLs(urls []string) []string {
	var productURLs []string
	// List of some products that are common on e-commerce websites
	patterns := []string{
		"product", "item", "electronics", "beauty", "mobile",
		"sneaker", "luxury", "toys", "home", "watch",
		"clothing", "shoes", "accessories", "furniture", "kitchen",
		"appliance", "gadget", "book", "stationery", "gaming",
		"sports", "fitness", "outdoor", "jewelry", "bag",
		"grocery", "health", "wellness", "cosmetics", "apparel",
		"kids", "pet", "hardware", "tool", "vehicle",
		"automobile", "bike", "camera", "audio", "video",
		"decor", "garden", "lighting", "office", "travel",
		"software", "subscription", "digital", "gift", "art",
		"craft", "diy", "hobby", "collectible", "memorabilia",
		"perfume", "skincare", "makeup", "haircare", "baby",
		"infant", "toddler", "maternity", "party", "events",
		"wedding", "costume", "seasonal", "festival", "outfit",
		"suit", "tie", "dress", "uniform", "lingerie",
		"bedding", "blanket", "pillow", "mattress", "curtain",
		"rug", "carpet", "cleaning", "storage", "organizer",
		"fitness equipment", "exercise", "protein", "supplement", "nutrition",
		"tea", "coffee", "beverage", "snack", "chocolate",
		"bakery", "meat", "seafood", "dairy", "frozen",
		"fruit", "vegetable", "organic", "herbal", "medicine",
		"eyewear", "sunglasses", "glasses", "contacts", "optical",
		"automotive", "tires", "engine", "battery", "parts",
		"tools", "construction", "paint", "plumbing", "heating",
		"cooling", "water", "filter", "security", "surveillance",
		"smart", "iot", "automation", "lighting", "energy",
		"solar", "generator", "power", "network", "router",
		"server", "storage", "accessory", "case", "cover",
		"headphone", "earbud", "speaker", "soundbar", "monitor",
		"laptop", "tablet", "desktop", "printer", "scanner",
		"office supply", "pen", "paper", "envelope", "binder",
		"file", "notebook", "planner", "calendar", "diary",
		"poster", "frame", "canvas", "sculpture", "vase",
		"antique", "vintage", "retro", "modern", "contemporary",
	}

	for _, url := range urls {
		for _, pattern := range patterns {
			if strings.Contains(url, pattern) {
				productURLs = append(productURLs, url)
				break
			}
		}
	}
	return productURLs
}

func Unique(urls []string) []string {
	seen := make(map[string]struct{})
	var result []string
	// abstract all the unique urls
	for _, url := range urls {
		if _, exists := seen[url]; !exists {
			seen[url] = struct{}{}
			result = append(result, url)
		}
	}
	return result
}
