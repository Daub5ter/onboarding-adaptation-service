import { useState } from "react";

export const SolveInstruction = () => {
    const [solved, setSolved] = useState();

    const handleShow = async () => {
        const payload = {
            action: "solve_instruction",
            users_instructions: {
                user_id: 1,
                instruction_id: 3,
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

        setSolved(data);
    }


    return (
        <>
            <button onClick={handleShow}>solve instruction</button>

            {solved && (
                <div>
                    <h2>Successful solved</h2>
                </div>

            )}
        </>
    )
}