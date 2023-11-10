const cors = require('cors');
const express = require('express');
const fileUpload = require('express-fileupload');
const port = 3000;

const app = express();

app.use(fileUpload({
    createParentPath: true,
}))

app.use(cors())

app.post('/upload', (req, res) => {
    if (!req.files) {
        console.log("no files uploaded")
        return res.status(400).json({msg: 'No file uploaded'});
    }

    const filesDir = encodeURI(Date.now())


    for (let i = 0; ; i++) {
        let file = req.files.file[i];
        if (!file) {
            break;
        }

        file.mv(`${__dirname}/images/${filesDir}/${file.name}`, err => {
            console.log(file.name)
            if (err) {
                console.error(err);
                return res.status(500).send(err);
            }
        });
    }

    res.json({
        filePath: `/images/${filesDir}`,
    });
});

app.listen(port, () => console.log(`Server Started in port ${port}`));