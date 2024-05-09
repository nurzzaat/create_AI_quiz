ClassicEditor.create(document.querySelector("#ckeditor"), {
    mediaEmbed: {
        previewsInData: true,
    },
}).catch((error) => {
    console.error(error)
})
