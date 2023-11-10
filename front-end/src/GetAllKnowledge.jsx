import {useState} from "react";

export const GetAllKnowledge = () => {
    const [showed, setShowed] = useState();

    const handleShow = async () => {
        const payload = {
            action: "get_all_knowledge",
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
            <button onClick={handleShow}>show all knowledge</button>

            {showed && (
                <div>
                    <h2>{showed.data[1].title}</h2>
                    <h2>{showed.data[1].description}</h2>
                    <h2>{showed.data[1].solved}</h2>
                </div>

            )}
        </>
    )

}