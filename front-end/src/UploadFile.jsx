import { useState, useRef } from "react";

const hostUrl = "http:\/\/localhost:3000/upload";

export const UploadFile = () => {
    const filePicker = useRef(null);
    const [selectedFile, setSelectedFile] = useState(null);
    const [uploaded, setUploaded] = useState();

    const handleChange = (event) => {
        setSelectedFile(event.target.files)
    };


    const handleUpload = async () => {
        if (!selectedFile) {
            alert("Please select a file");
            return;
        }

        const formData = new FormData();

        for (let i = 0; ; i++) {
            if (selectedFile[i] === undefined) {
                break
            }
            formData.append('file', selectedFile[i]);
        }

        const res = await fetch(hostUrl, {
            method: 'POST',
            body: formData,
        });
        const data = await res.json();

        setUploaded(data);
    };

    const handlePick = () => {
        filePicker.current.click();
    }

    return (
        <>
            <button onClick={handlePick}>Pick files</button>
            <input
                className="hidden"
                type="file"
                ref = {filePicker}
                multiple
                onChange={handleChange}
                accept="image/*,.png,.jpg,.jpeg,.gif,.web,"
            />

            <button onClick={handleUpload}>Upload now</button>

            {selectedFile && (
                <ul>
                    <li>Name: {selectedFile.name}</li>
                    <li>Type: {selectedFile.type}</li>
                    <li>Size: {selectedFile.size}</li>
                </ul>
            )}

            {uploaded && (
                <div>
                    <h2>{uploaded.filePath}</h2>
                    <img alt='' src={uploaded.filePath} width="200"/>
                </div>
            )}
        </>
    );
};