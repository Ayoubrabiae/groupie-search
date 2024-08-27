// Search Button Functionality
const searchHolder = document.querySelector(".search-btn")
const searchBtn = document.querySelector(".search-btn i")
const searchInp = document.querySelector(".search-btn input")

if (searchBtn) {
    searchBtn.addEventListener("click", () => {
        searchHolder.classList.toggle("show")
        searchBtn.classList.toggle("bx-x")
        searchInp.focus()
    })
}

// Search Functionality
const search = async(value) => {
    const res = await fetch(`/suggest-search?q=${value}`)

    return await res.json()
}

if (searchInp) {
    searchInp.addEventListener("input", (e) => {
        const val = e.currentTarget.value

        if (val == "") {
            return
        }

        const data = search(val)

        console.log(data)
    })
}