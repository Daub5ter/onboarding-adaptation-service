import { useState } from "react";

export const AddUsersKnowledge = (userID, knowledgeID) => {
    const [added, setAdded] = useState();

    const handleShow = async (userID, knowledgeID) => {
        const payload = {
            action: "add_users_knowledge",
            users_known: {
                user_id: userID,
                knowledge_id: knowledgeID,
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

    return handleShow(userID, knowledgeID);
}
    /*return (
        <>
            <button onClick={handleShow}>add users knowledge</button>

            {added && (
                <div>
                    <h2>Successful added</h2>
                </div>

            )}
        </>
    )
}*/