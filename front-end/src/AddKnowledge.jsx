import { useState } from "react";

export const AddKnowledge = () => {
    const [added, setAdded] = useState();

    const handleShow = async () => {
        const payload = {
            action: "add_knowledge",
            known: {
                title: "ddada",
                description: "desc",
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
            <button onClick={handleShow}>add knowledge</button>

            {added && (
                <div>
                    <h2>Successful added</h2>
                </div>

            )}
        </>
    )
}