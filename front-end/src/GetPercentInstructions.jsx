import {useState} from "react";

export const GetPercentInstructions = () => {
    const [showed, setShowed] = useState();

    const handleShow = async () => {
        const payload = {
            action: "get_percent_instructions",
            id: {
                id: 1,
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

        setShowed(data);
    }


    return (
        <>
            <button onClick={handleShow}>get percent instructions</button>

            {showed && (
                <div>
                    <h2>{showed.data}</h2>
                </div>

            )}
        </>
    )

}