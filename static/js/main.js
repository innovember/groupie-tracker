

function setData(data) {


}


function renderArtist(data) {
    // <div class="w3-quarter w3-container w3-margin-bottom">
    //     <img src="https://groupietrackers.herokuapp.com/api/images/queen.jpeg" alt="Norway" style="width:100%"
    //     class="w3-hover-opacity">
    //   <div class="w3-container w3-white">
    //     <p><b>Lorem Ipsum</b></p>
    //     <p>Praesent tincidunt sed tellus ut rutrum. Sed vitae justo condimentum, porta lectus vitae, ultricies congue
    //       gravida diam non fringilla.</p>
    //     <button type="submit" class="w3-button w3-black w3-margin-bottom w3-block">
    //       <i class="fa fa-paper-plane w3-margin-right"></i>Send Message</button>
    //   </div>
    // </div>


    var img = '<img src="' + data.image + '" alt="Norway" style="width:100%" class="w3-hover-opacity">'

    var detailsCont = document.createElement("div");
    detailsCont.className = "w3-container w3-white"
    detailsCont.append('<p class="w3-large"><b>' + data.name + '</b></p>')
    detailsCont.append('<p><span class="w3-text-grey">Formed</span><br>' + data.members.join(', ') + "</p>")
    detailsCont.append('<p><span class="w3-text-grey">Formed</span><br>' + data.creationDate + "</p>")
    detailsCont.append('<p><span class="w3-text-grey">First Album</span><br>' + data.firstAlbum + "</p>")
    detailsCont.append('<button type="submit" onClick="open(' + data.id + ')" class="w3-button w3-black w3-margin-bottom w3-block">  ' +
        '<i class="fa fa-paper-plane w3-margin-right"></i>Send Message</button>')

    var card = document.createElement("div");
    node.className = "w3-quarter w3-container w3-margin-bottom"
    card.appendChild(img)
    card.appendChild(detailsCont)

    return card
}
