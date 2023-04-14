$(document).ready(function () {
    $('#upload-button').click(function () {
        var formData = new FormData($('#upload-form')[0]);
        $.ajax({
            url: "/upload",
            type: "POST",
            data: formData,
            processData: false,
            contentType: false,
            success: function (response) {
                console.log(response);
            }
        });
    });
});


