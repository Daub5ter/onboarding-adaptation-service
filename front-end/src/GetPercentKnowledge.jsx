import {useState} from "react";

export const GetPercentKnowledge = (userID) => {

    const handleShow = async (userID) => {
        const payload = {
            action: "get_percent_knowledge",
            id: {
                id: userID,
            }
        }

        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const response = await fetch("http:\/\/localhost:8080/handle", {
            method: 'POST',
            body: JSON.stringify(payload),
            headers: headers,
        });

        return await response.json();
    }

    return handleShow(userID);
}

/*    return (
        <>
            <button onClick={handleShow}>get percent knowledge</button>

            {showed && (
                <div>
                    <h2>{showed.data}</h2>
                </div>

            )}
        </>
    )

}*/