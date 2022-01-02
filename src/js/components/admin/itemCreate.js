import React, { useState, useEffect } from "react";
import ReactDOM from "react-dom";
import axios from 'axios';

import ItemImage from "./itemImage";


const ItemCreate = () =>  {

  // let csrf = document.getElementsByName("csrf_token")[0].value
  // console.log(csrf)
  var category = document.getElementById("category").value;
  var res = (categories !== null) ? JSON.parse(category) : [{ID:'', name:''}];

  const [data, setData] = useState([])
  const [image, setImage] = useState("")
  const [imageIds, setImageIds] = useState([])
  const [imagePath, setImagePath] = useState([])
  const [checked, setChecked] = useState(false);
  const [categories, setCategories] = useState(res[0]);
  const [previews, setPreviews] = useState([])
  const [postImage, setPostImage] = useState()

  const [hasError, setError] = React.useState(false);

  if (image !== '') {
    imageIds[image.name] = image.ID
    imagePath.push(image.file_path)
    setData({...data, imageIds})
    setPreviews(imagePath)
    setImage('')
  }

  const postItem = (data) => {
    if (data.length ===  0) {
      return false
    }
    var itemData = appendData(data);
    console.log(data)

    axios.post( '/admin/api/post-item', itemData, {
            headers: {
              'X-Requested-With': 'XMLHttpRequest',
              'content-type': 'multipart/form-data',
              // "X-CSRF-Token": csrf
            },
        }
    )
    .then(response => {
        console.log(response.data);
    })
    .catch(err => {
        console.log(err);
        setError(true);
    });
  };

  if (hasError) {
    return <p>Sorry, Sign up failed!</p>;
  }

  const appendData = (data) => {
    const itemData = new FormData()

    itemData.append("name", data.name)
    console.log(data)
    
    if (data.detail === undefined) data.detail = '';
    itemData.append("detail", data.detail)

    if (data.normal_price === undefined) data.normal_price = 0;
    itemData.append("normal_price", data.normal_price)

    if (data.special_price === undefined) data.special_price = 0;
    itemData.append("special_price", data.special_price)
    
    itemData.append("stock", parseInt(data.stock))
    // itemData.append("add_point", data.add_point)

    if (data.display_flg === undefined) data.display_flg = 0;
    itemData.append("display_flg", data.display_flg)

    if (isNaN(data.category_id)) {
      data.category_id = null;
    } else {
      data.category_id = parseInt(data.category_id);
    }
    itemData.append("category_id", data.category_id)
    itemData.append("item_image_id[]", data.imageIds)
    return itemData;
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