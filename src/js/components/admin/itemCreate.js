import React, { useState, useEffect } from "react";
import ReactDOM from "react-dom";
import axios from 'axios';

const ItemCreate = () =>  {

  const [data, setData] = useState([])
  const [image, setImage] = useState("")
  const [checked, setChecked] = useState(false);

  const onSubmit = (data) => {
    console.log(data)

    const itemData = new FormData()
    itemData.append("name", data.name)
    itemData.append("detail", data.detail)
    itemData.append("normal_price", data.normal_price)
    itemData.append("special_price", data.special_price)
    itemData.append("stock", data.stock)
    
    if (data.display_flg === undefined) {
        data.display_flg = 0;
    }
    itemData.append("display_flg", data.display_flg)
    itemData.append("image", fileInput.current.files[0])

    axios.post( '/admin/item/store', itemData, {
            headers: {
              'X-Requested-With': 'XMLHttpRequest',
              'content-type': 'multipart/form-data',
            },
        }
    )
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
  };


  const fileInput = React.createRef()

  return (
      <>
        <p>
          <input 
            type="file" 
            multiple="multiple"
            name="image" 
            ref={fileInput}
            onChange={handleChange}
            accept="image/*"
          />
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
        <label>特別価格</label>
          <input 
            type="text"
            name="special_price"
            // value={data.detail}
            required="required"
            onChange={handleChange}
          />
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

          <button type="submit" onClick={() => onSubmit(data)} />登録する
      </>
  )
}

const ItemForm = document.getElementById("item_create");

ReactDOM.render(
  <ItemCreate/>, ItemForm
);

// export default ItemEdit