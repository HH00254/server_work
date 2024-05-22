document.addEventListener("DOMContentLoaded", load)

function load() {
    let userButton = document.getElementById("form-btn");
    let postButton = document.getElementById("btn-private-key-post"); 

    userButton.addEventListener("click",  async (event) => {

        if (event) {
            const [userName , userPassword] =  getFormData();
            const pubKey =  await getPublicKey();
            const encryptPubKey = await encryptClientPassword(userPassword, pubKey);

            // localStorage.setItem("password" ,encryptPubKey)
            
            setCookie(userName, encryptPubKey, 1);
        }

        event.preventDefault();
    });

    postButton.addEventListener("click", () => {
        let key = getCookieValue("uidPassword");

    });

   
}

async function getPublicKey() {
    const resquestURL = "/pgkey";

    try {
        const response     = await fetch(resquestURL);
        const dataFormated = await response.json();
      
        if (response.status === 200 &&
            dataFormated['public_key'] !== undefined &&
            dataFormated['public_key'].length > 0) {
            
            const publicKey = dataFormated["public_key"];
        
            return publicKey;

        } else {
            console.log("error");

        }

    } catch (error) {
        print('Fetch Error: ' + `${error.name.message}`);
        
    }
}

async function encryptClientPassword(password, publicKeyArmored) {

    const publicKey = await openpgp.readKey({ armoredKey: publicKeyArmored });

    const encrypted = await openpgp.encrypt({
        message: await openpgp.createMessage({ text: password }), 
        encryptionKeys: publicKey,
    });

    return encrypted;
}

function setCookie(cname, cvalue, exdays) {
    const date = new Date();
    date.setTime(date.getTime() + (exdays * 24*60*60*1000));

    let expires = "expires="+ date.toUTCString();
    document.cookie = "uidPassword" + '=' + encodeURIComponent(cvalue) + expires + "; Secure path=/";

}

function getCookieValue(keyName) {

    const value = document.cookie
        .split("; ")
        .find((row) => row.startsWith(keyName))
        ?.split("=")[1]
        .split("expires")[0];

        return decodeURIComponent(value);
}

function getFormData() {
    const name = document.getElementById("uid").value;
    const password = document.getElementById("uid-password").value;

    return [name , password]
}


