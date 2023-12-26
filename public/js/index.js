document.addEventListener('DOMContentLoaded', function () {
    // Find all elements with the 'image-preview' attribute
    var previewElements = document.querySelectorAll('[image-preview]');

    // Attach event listeners to each element
    previewElements.forEach(function (element) {
        // Get the value of the 'image-preview' attribute
        var previewId = element.getAttribute('image-preview');

        // Find the corresponding image element by ID
        var imageElement = document.getElementById(previewId);

        // Attach the change event listener to the file input
        element.addEventListener('change', function () {
            // Display the selected image in the preview
            displayImagePreview(this, imageElement);
        });
    });

    // Function to display the image preview
    function displayImagePreview(input, imageElement) {
        if (input.files && input.files[0]) {
            var reader = new FileReader();

            reader.onload = function (e) {
                // Set the source of the image element to the selected image
                imageElement.src = e.target.result;
            };

            // Read the selected file as a data URL
            reader.readAsDataURL(input.files[0]);
        }
    }
});
