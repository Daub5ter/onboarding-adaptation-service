<?php
    if($_FILES['filename']['size'] > 3*1024*1024) {
        exit('Not more 3 MB');
    }

    if(move_uploaded_file($_FILES['filename']['tmp_name'], 'temp/'.$_FILES['filename']['name'])) {
        echo 'File uploaded';
        echo 'File was uploaded to - ' . $_FILES['filename']['tmp_name'] . '<br>';
    } else {
        echo 'Error uploading';
    }
?>