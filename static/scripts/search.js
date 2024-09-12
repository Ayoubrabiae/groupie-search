// Search Button Functionality
const searchHolder = document.querySelector(".search-btn")
const searchBtn = document.querySelector(".search-btn i")
const searchInp = document.querySelector(".search-btn input")
const suggestions = document.querySelector(".search-holder .suggestions")


if (searchBtn) {
    searchBtn.addEventListener("click", () => {
        if (searchBtn.classList.contains("bx-x")) {
            searchInp.value = ""
            suggestions.innerHTML = ""
        } else {
            searchInp.focus()
        }
        searchHolder.classList.toggle("show")
        searchBtn.classList.toggle("bx-x")
    })
}

// Search Functionality
const search = async(value) => {
    const res = await fetch(`/suggest-search?q=${value}`)

    console.log(value+"|")

    return await res.json()
}

const creatSuggesition = (value, type, id) => {
    const suggestion = document.createElement("a")
    suggestion.classList.add("sug")
    suggestion.href = `/artists/${id}`

    const valueElement = document.createElement("span")
    valueElement.classList.add("name")
    valueElement.textContent = value
    
    const typeElement = document.createElement("span")
    typeElement.classList.add("type")
    typeElement.textContent = type
    suggestion.append(valueElement, typeElement)
    
    return suggestion
}

if (searchInp) {
    searchInp.addEventListener("input", async (e) => {
        suggestions.innerHTML = ""
        let val = e.currentTarget.value

        console.log(val.split("").some(e => e != " "))
        if (val == "" || !val.split("").some(e => e != " ")) {
            return
        }
        val = val.split(" ").filter(e => e != "").join(" ")

        const res = await search(val)

        if (!res.length) {
            const suggestion = document.createElement("p")
            suggestion.classList.add("sug")
            suggestion.textContent = `${val} Not Found`

            suggestions.append(suggestion)
            return
        }


        for (let r of res) {
                if (r.Kind === "band") {
                    suggestions.append(creatSuggesition(r.Value, r.Kind, r.Id))
                } else {
                    suggestions.append(creatSuggesition(r.Value, `${r.Kind} of ${r.Name}`, r.Id))
                }
        }

    })
}