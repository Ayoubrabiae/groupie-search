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

    return await res.json()
}

const creatSuggesition = (value, complete, type, id) => {
    const suggestion = document.createElement("a")
    suggestion.classList.add("sug")
    suggestion.href = `/artists/${id}`

    const valueElement = document.createElement("span")
    valueElement.classList.add("name")
    valueElement.textContent = value
    
    const completeElement = document.createElement("span")
    completeElement.classList.add("complete")
    completeElement.textContent = complete
    valueElement.append(completeElement)
    
    const typeElement = document.createElement("span")
    typeElement.classList.add("type")
    typeElement.textContent = type
    suggestion.append(valueElement, typeElement)
    
    return suggestion
}

if (searchInp) {
    searchInp.addEventListener("input", async (e) => {
        suggestions.innerHTML = ""
        const val = e.currentTarget.value

        if (val == "") {
            return
        }

        const data = await search(val)

        if (!data.length) {
            const suggestion = document.createElement("p")
            suggestion.classList.add("sug")
            suggestion.textContent = `${val} Not Found`

            suggestions.append(suggestion)
            return
        }

        for (let d of data) {
            suggestions.append(creatSuggesition(val, d.Value, d.Kind, d.Id))
        }

    })
}