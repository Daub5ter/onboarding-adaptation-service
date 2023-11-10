import {useState} from "react";

export const GetAllKnowledge = (id) => {
    const handleShow = async (id) => {
        const payload = {
            action: "get_all_knowledge",
            id: {
                id: id,
            }
        }

        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const response = await fetch("http:\/\/localhost:8080/handle", {
            method: 'POST',
            body: JSON.stringify(payload),
            headers: headers,
        });

        return await response.json()
    }

    return handleShow(id);

    /*return (
        <>
            <button onClick={handleShow}>show all knowledge</button>

            {showed && (
                <div>
                    <h2>{showed.data[1].title}</h2>
                    <h2>{showed.data[1].description}</h2>
                </div>

            )}
        </>
    )*/

}