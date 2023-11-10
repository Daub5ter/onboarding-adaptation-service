import { useState } from "react";

export const AddUsersKnowledge = () => {
    const [added, setAdded] = useState();

    const handleShow = async () => {
        const payload = {
            action: "add_users_knowledge",
            users_known: {
                user_id: 1,
                knowledge_id: 1,
            }
        }

        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const response = await fetch("http:\/\/localhost:8080/handle", {
            method: 'POST',
            body: JSON.stringify(payload),
            headers: headers,
        });

        const data = await response.json();

        setAdded(data);
    }


    return (
        <>
            <button onClick={handleShow}>add users knowledge</button>

            {added && (
                <div>
                    <h2>Successful added</h2>
                </div>

            )}
        </>
    )
}