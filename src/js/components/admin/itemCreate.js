import React, { useState, useEffect, forwardRef } from "react";
import ReactDOM from "react-dom";
import axios from 'axios';

import ItemImage from "./itemImage";
import ItemValidation from "../../lib/itemValidation";

const ItemCreate = () =>  {

  // let csrf = document.getElementsByName("csrf_token")[0].value
  // console.log(csrf)
  var category = document.getElementById("category").value;
  var res = (categories !== null) ? JSON.parse(category) : [{ID:'', name:''}];
  
  const numberRegExp = /^[0-9]+$/

  const [data, setData] = useState([])
  const [image, setImage] = useState("")
  const [item_image_id, setImageIds] = useState([])
  const [imagePath, setImagePath] = useState([])
  const [checked, setChecked] = useState(false);
  const [categories, setCategories] = useState(res[0]);
  const [previews, setPreviews] = useState([])

  const [error, setError] = useState({'valid': {'key': [], 'message': []}});
  
  axios.defaults.headers.common = {
    'Content-Type': 'application/json',
    'X-Requested-With': 'XMLHttpRequest',
    // "X-CSRF-Token": csrfToken
  }

  if (image !== '') {
    item_image_id[image.name] = {'id': image.ID}
    imagePath.push(image.file_path)
    setData({...data, item_image_id})
    setPreviews(imagePath)
    setImage('')
  }

  const postItem = (data) => {
    console.log(data)
    
    const [valid, hasError] = ItemValidation(data)
    if (hasError === true) return setError({...error, valid})

    axios.post('/admin/api/post-item', data)
    .then(response => {
        console.log(response.data);
    })
    .catch(err => {
        console.log(err);
    });
  };


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
            // value={setData.name}
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
            // value={data.detail}
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
          >
            <option value="" hidden>選択してください</option>
            {categories.map((v, i) =>
                <option  value={v.ID} key={i}>{v.name}</option>
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