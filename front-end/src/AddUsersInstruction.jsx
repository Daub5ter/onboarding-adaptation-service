import { useState} from "react";

export const AddUsersInstruction = (userID, instructionID) => {
    const [added, setAdded] = useState();

    const handleShow = async (userID, instructionID) => {
        const payload = {
            action: "add_users_instruction",
            users_instructions: {
                user_id: userID,
                instruction_id: instructionID,
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

    return handleShow(userID, instructionID);
}

    /*return (
        <>
            <button onClick={handleShow}>add users instruction</button>

            {added && (
                <div>
                    <h2>Successful added</h2>
                </div>

            )}
        </>
    )
}*/