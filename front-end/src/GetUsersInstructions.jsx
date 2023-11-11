import {useState} from "react";

export const GetUsersInstructions = (id) => {
    const [showed, setShowed] = useState();

    const handleShow = async (id) => {
        const payload = {
            action: "get_users_instructions",
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

        return await response.json();
    }

    return handleShow(id);

}