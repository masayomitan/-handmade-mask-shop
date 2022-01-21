import React, { useState, useEffect } from "react";
import axios from 'axios';

const ItemImage = (props) => {
  
  const [modalShow, setModalShow] = useState(false)
  const [selectedId, setSelectedId] = useState('')
  const [images, setImages] = useState([])
  const [confirm, setConfirm] = useState(false)
  
  const fileInput = React.createRef()
  const [num, setNum] = useState(1)  //count fileupload
  
  axios.defaults.headers.common = {
    'Content-Type': 'application/json',
    'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }

  const toggleModal = (val) => {
    setModalShow(val)
    if (val === true) {
      getItemImages()
    }
  }

  
  // const handlePreview = (path) => {
  //   if (path !== null) {
  //     previews.push(path)
  //     setPreview(previews)
  //   } else {
  //     return setPreviews(null)
  //   }
  // }


  const onFileChange = (e) => {
    const files = e.target.files
    const name = e.target.name
    if (files.length > 0) {
      var file = files[0]
      saveImage(file, name)
      // var reader = new FileReader()
      // reader.onload = (e) => {
      //   return handlePreview(e.target.result)
      // }
      // reader.readAsDataURL(file)
    }
  }

  const saveImage = (file, name) => {
    const param = new FormData();
    param.append('file', file)
    axios.post('/admin/api/post-item-image', param, {
      headers: {
        'content-type': 'multipart/form-data',
        // "X-CSRF-Token": csrf
      },
    })
    .then(response => {
      let value = response.data
      value.name = name
      props.setImage(value)
      if (5 > num) setNum(num + 1)
      return response.data
    })
    .catch(err => {
        console.log(err);
    });
  }

  const getItemImages = async() => {
    try {
      const result = await axios.get('/admin/api/get-item-images')
      let value = result.data
      setImages(value)
    } catch (e) {
      dispatchEvent(e)
      if (e.response && e.response.status === 400) {
        console.log('400 Error!!')
      }
    }
  }

  const confirmSelect = (val, e) => {
    val.name = e.target.name
    setConfirm(() => confirmButton(val))
    setSelectedId(val.ID)
  }

  const setImageData = (val, bool) => {
    if (bool === true) {
      props.setImage(val)
      if (5 > num) setNum(num + 1)
    }
      setModalShow(false)
      setSelectedId("")
  }

  const confirmButton = (val) => {
    return (
      <>
        <button 
          onClick={() => setImageData(val, true)}
        >
          選択
        </button>

        <button
          onClick={() => setImageData(val, false)}
        >
          戻る
        </button>
      </>
    )
  }

  return (
    <>
      {[...Array(num)].map((e, i) => (
        <p key={i}>
          <input 
            type="file" 
            multiple="multiple"
            name={"imageId" + num}
            ref={fileInput}
            onChange={(e) => onFileChange(e)}
            accept="image/*"
          />

          <button
            onClick={() => toggleModal(true)}
            name={"imageId" + num}
          >
          フォルダから選択
          </button>
        </p>
      ))}
        
        <div>
          {modalShow &&
            <div className="flex items-center justify-center fixed left-0 bottom-0 w-full h-full bg-gray-800">
              <div className="bg-white rounded-lg ">
                <div className="flex flex-col items-start p-4">
                  <div className="flex items-center w-full">
                    {images.length > 0 ?
                    
                      images.map((v, i) => {
                        return (
                          <label key={i}>
                            <img
                              name={"imageId" + num}
                              className={(v.ID === selectedId) ? "opacity-50" : ""}  
                              src={v.file_path} 
                              onClick={(e) => confirmSelect(v, e)} />
                          </ label>
                        )
                      })
                      :
                      <>
                        <span>画像がまだありません</span>
                        <button
                          onClick={() => setImageData(null, false)}
                        >
                          戻る
                        </button>
                      </>
                    }
                  </div>
                </div>
                {confirm}
              </div>
            </div>
          }
        </div>
    </>
  )

}

export default ItemImage