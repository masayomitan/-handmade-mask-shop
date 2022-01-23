import React, { useState, useEffect } from "react";
import ReactDOM from "react-dom";
import axios from 'axios';

import ItemImage from "./itemImage";
import ItemValidation from "../../lib/itemValidation";

const ItemCreate = () =>  {
  axios.defaults.headers.common = {
    'content-Type': 'multipart/form-data',
    'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }
  // let csrf = document.getElementsByName("csrf_token")[0].value
  // console.log(csrf)
  var category = document.getElementById("category").value;
  var res = (categories !== null) ? JSON.parse(category) : [{ID:'', name:''}];

  const numberRegExp = /^[0-9]+$/

  const [data, setData] = useState([])

  const itemData = document.getElementById("item_edit")
  const [editData, setEditData] = useState((itemData !== null) ? JSON.parse(itemData.value) : null)
  
  console.log(editData[0].ItemImages)

  const [image, setImage] = useState("")
  const [imageIds, setImageIds] = useState([])
  const [imagePath, setImagePath] = useState([])
  
  const [checked, setChecked] = useState(false);
  const [categories, setCategories] = useState(res[0]);
  const [previews, setPreviews] = useState([])


  const match = location.href.match(/[0-9]+/)
  const id = (match !== null) ? match[0] : null
  const endPoint = (id !== null)  ? "/admin/api/update-item/" + id : "/admin/api/post-item"
  
  const [error, setError] = useState({'valid': {'key': [], 'message': []}});


  if (image !== '') {
    imageIds[image.name] = image.ID
    imagePath.push(image.file_path)
    setData({...data, imageIds})
    setPreviews(imagePath)
    setImage('')
  }

  const postItem = (data) => {
    
    const [valid, hasError] = ItemValidation(data)
    if (hasError === true) return setError({...error, valid})
    
    const setParams = appendData(data)

    axios.post(endPoint, setParams)
    .then(response => {
        console.log(response.data);
    })
    .catch(err => {
        console.log(err);
    });
  };

  const appendData = (data) => {
    const param = new FormData();
    param.append("name", data.name)
    param.append("detail", data.detail)
    if (data.normal_price === undefined) data.normal_price = 0;
    param.append("normal_price", data.normal_price)
    if (data.special_price === undefined) data.special_price = 0;
    param.append("special_price", data.special_price)
    
    
    if (data.category_id !== undefined) {
      if (isNaN(data.category_id)) {
        data.category_id = null;
      } else {
        data.category_id = parseInt(data.category_id);
      }
    }
    param.append("category_id", data.category_id)

    if (data.imageIds !== undefined) {
      console.log(data.imageIds)
      let num = Object.keys(data.imageIds).length
      for (let i = 0; i < num; i++) {
        let sum = (i+1)
        param.append("imageId" + sum, data.imageIds["imageId" + sum])
      }
    }
    return param
  }


  const handleChange = (e) => {
    const t = e.target;
    const name = t.name;
    const value = t.type === "file" ? t.files[0] : t.value;
    setData({...data, [name]: value})
  }


  const handleCheck = (e) => {
    const t = e.target;
    const name = t.name;
    const value = t.checked === true ? 1 : 0;
    setData({...data, [name]: value})
    return setChecked(value);
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

        <ItemImage
          setImage={setImage}
        />

        <p className="text-red-500">
          {error.valid.key.name === true && (error.valid.message.name)}
        </p>
        <p>
          <label>商品名</label>
          <input 
            type="text"
            name="name"
            defaultValue={(editData !== null) ? editData[0].Name: ''}
            required="required"
            onChange={handleChange}
          />
        </p>

        <p className="text-red-500">
          {error.valid.key.detail === true && (error.valid.message.detail)}
        </p>

        <p>
          <label>詳細</label>
          <input 
            type="textarea"
            name="detail"
            defaultValue={(editData !== null) ? editData[0].Detail : ''}
            required="required"
            onChange={handleChange}
          />
        </p>
        
        <p>
          <label>通常価格</label>
          <input 
            type="text"
            name="normal_price"
            // value={data.detail}
            required="required"
            onChange={handleChange}
          />
        </p>
        <p>
          <label>特別価格</label>
          <input 
            type="text"
            name="special_price"
            // value={data.detail}
            required="required"
            onChange={handleChange}
          />
        </p>
        <p>
          <label>在庫</label>
          <input 
            type="number"
            name="stock"
            // value={data.detail}
            required="required"
            onChange={handleChange}
          />
        </p>
        <p>
          <label>公開チェック</label>
          <input 
            type="checkbox"
            name="display_flg"
            checked={checked}
            defaultValue={checked}
            {...checked ? true : false}
            onChange={handleCheck}
          />
        </p>

        <p className="text-red-500">
          {error.valid.key.category === true && (error.valid.message.category)}
        </p>
        <p>
          <label>カテゴリー</label>
          <select
              name="category_id"
              onChange={handleChange}
              defaultValue={(editData !== null) && editData[0].Category.ID}
          >
            <option value="" hidden>選択してください</option>
            {categories.map((v, i) =>
                <option
                    value={v.ID} key={i}>{v.name}
                </option>
              )
          }
          </select>
        </p>
        <button type="submit" onClick={() => postItem(data)} >登録す</button>
      </>
  )
}

if (document.getElementById("item_create")) {
  ReactDOM.render(
    <ItemCreate />, document.getElementById("item_create")
  );
}