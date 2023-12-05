function fetchUserData(sessionToken, setLoad) {
    let payloadSession = {
        action: "authenticate_user_session",
        session: {
            session_token: sessionToken,
        }
    }

    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    return fetch("http:\/\/localhost:8080/handle", {
        method: 'POST',
        body: JSON.stringify(payloadSession),
        headers: headers,
    })
        .then(response => response.json())
        .then(dataSession => {
            if (dataSession.error) {
                localStorage.removeItem('session_token');
                console.log("Session is not valid");
                setLoad(true);
                return dataSession;
            } else {
                let payload = {
                    action: "get_user_by_email",
                    email: {
                        email: dataSession.data,
                    }
                }

                fetch("http:\/\/localhost:8080/handle", {
                    method: 'POST',
                    body: JSON.stringify(payload),
                    headers: headers,
                })
                    .then((response) => response.json())
                    .then((data) => {
                        if (data.error) {
                            console.log("Error to fetch auth user function")
                        } else {
                            dataSession = data
                        }
                    })
                    .catch((error) => {
                        console.log(error)
                    })
            }
            return dataSession
        })
        .catch(error => console.error(error));
}

function LoadSession(sessionToken, setLoad, setIsLoggedIn, setUsername) {
    fetchUserData(sessionToken, setLoad)
        .then(data => {
            if (data.error !== true) {
                setIsLoggedIn(true);
                setUsername(data.data);
                setLoad(true);
            }
        })
        .catch(error => {
            console.error(error)
        });
}

export default LoadSession;
