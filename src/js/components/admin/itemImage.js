import React, { useState, useEffect } from "react";
import axios from 'axios';

const ItemImage = (props) => {
  
  const [image, setImage] = useState('')
  const [images, setImages] = useState([])
  const [preview, setPreview] = useState(null)
  const [previews, setPreviews] = useState([])
  const fileInput = React.createRef()
  const n = 5  //count fileupload


  const handlePreview = (path) => {
    if (path !== null) {
      previews.push(path)
      setPreviews(previews)
    } else {
      return setPreviews(null)
    }
  }


  const onFileChange = (e) => {
    const files = e.target.files
    if (files.length > 0) {
      var file = files[0]
      var reader = new FileReader()
      reader.onload = (e) => {
        return handlePreview(e.target.result)
      }
      reader.readAsDataURL(file)
      saveImage(file)
    }
  }

  const saveImage = (file) => {
    const param = new FormData();
    param.append('file', file)
    axios.post('/admin/api/post-item-image', param, {
      headers: {
        'X-Requested-With': 'XMLHttpRequest',
        'content-type': 'multipart/form-data',
        // "X-CSRF-Token": csrf
      },
    })
    .then(response => {
      const value = response.data
      props.setImage(value)
      return response.data
    })
    .catch(err => {
        console.log(err);
    });
  }


  return (
    <>
      {previews.map((path, i) => (
        <div key={i} className="w-20">
          <p>
            <img src={path} />
          </p>
        </div>
      ))}

      {[...Array(n)].map((e, i) => (
        <p key={i}>
          <input 
            type="file" 
            multiple="multiple"
            name="image" 
            ref={fileInput}
            onChange={(e) => onFileChange(e)}
            accept="image/*"
          />
        </p>
      ))}
    </>
  )

}

export default ItemImage