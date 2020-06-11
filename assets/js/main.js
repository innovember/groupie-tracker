function setData(data) {
    mainContainer.innerHTML = ""
    for (i = 0; i < data.length; i++) {
        // var q = Math.floor(y/x);
        if (i % 4 == 0) {
            var row = document.createElement("div")
            row.className = "w3-row-padding"
            mainContainer.appendChild(row)
            // row.innerHTML = ""
        }
        card = renderArtist(data[i])
        row.appendChild(card)
    }
    hideLoader()
}

function renderArtist(data) {
    // <div class="w3-quarter w3-container w3-margin-bottom">
    //     <img src="https://groupietrackers.herokuapp.com/api/images/queen.jpeg" alt="Norway" style="width:100%"
    //     class="w3-hover-opacity">
    //   <div class="w3-container w3-white">
    //     <p class="p"><b>Lorem Ipsum</b></p>
    //     <p>Praesent tincidunt sed tellus ut rutrum. Sed vitae justo condimentum, porta lectus vitae, ultricies congue
    //       gravida diam non fringilla.</p>
    //     <button type="submit" class="w3-button w3-black w3-margin-bottom w3-block">
    //       <i class="fa fa-paper-plane w3-margin-right"></i>Send Message</button>
    //   </div>
    // </div>


    var img = '<img src="' + data.image + '" alt="Norway" style="width:100%" class="w3-hover-opacity">'

    var detailsCont = document.createElement("div");
    detailsCont.className = "w3-container w3-white"
    detailsCont.innerHTML =
        '<p class="w3-large"><b>' + data.name + '</b></p>' +
        '<p class="p"><span class="w3-text-grey">Members</span><br>' + data.members.join(', ') + "</p>" +
        '<p class="p"><span class="w3-text-grey">Formed</span><br>' + data.creationDate + "</p>" +
        '<p class="p"><span class="w3-text-grey">First Album</span><br>' + data.firstAlbum + "</p>" +
        '<a href="http://127.0.0.1:8181/artist_info/' + data.id + '" class="w3-button w3-black w3-margin-bottom w3-block" target="_blank">  ' +
        '<i class="fa fa-eye w3-margin-right"></i>View</button>'

    var card = document.createElement("div");
    card.className = "w3-quarter w3-container w3-margin-bottom"
    // card.append(img)
    card.innerHTML = img
    card.appendChild(detailsCont)

    return card
}


function open(id) {
    location.replace("127.0.0.1:8181/artist/" + id)
}


function showLoader() {
    $("#loader").css('visibility', 'visible');
    $("#mainContainer").css('visibility', 'hidden');
}

function hideLoader() {
    $("#loader").css('visibility', 'hidden');
    $("#mainContainer").css('visibility', 'visible');
}