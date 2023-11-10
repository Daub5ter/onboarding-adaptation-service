import { useState} from "react";

export const AddUsersInstruction = () => {
    const [added, setAdded] = useState();

    const handleShow = async () => {
        const payload = {
            action: "add_users_instruction",
            users_instructions: {
                user_id: 2,
                instruction_id: 1,
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
            <button onClick={handleShow}>add users instruction</button>

            {added && (
                <div>
                    <h2>Successful added</h2>
                </div>

            )}
        </>
    )
}