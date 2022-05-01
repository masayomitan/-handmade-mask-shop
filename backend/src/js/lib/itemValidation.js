
const ItemValidation = (data) =>  {

  const valid = {'key': [], 'message': []}
  let hasError = false


  if (data.length ===  0) {
    hasError = true
    valid.key.all = true
    valid.message.all = '商品情報が未入力です'
  }
  console.log(data.category_id)
  if (!data.name) {
    hasError = true
    valid.key.name = true
    valid.message.name = '商品名を入力してください'
  } else if (data.name.length > 50) {
    hasError = true
    valid.key.name = true
    valid.message.name = '商品名は50文字以内で入力してください'
  }

  if (!data.detail) {
    hasError = true
    valid.key.detail = true
    valid.message.detail = '詳細を入力してください'
  } else if (data.detail.length > 2000){
    hasError = true
    valid.key.detail = true
    valid.message.detail = '詳細は2000文字以内で入力してください'
  }
  // }
  // if (data.normal_price === undefined) data.normal_price = 0;
  // itemData.append("normal_price", data.normal_price)

  // if (data.special_price === undefined) data.special_price = 0;
  // itemData.append("special_price", data.special_price)
  
  // itemData.append("stock", parseInt(data.stock))
  // // itemData.append("add_point", data.add_point)
  // if (data.display_flg !== undefined) data.display_flg = null;
  
  if (data.category_id !== undefined) {
    if (isNaN(data.category_id)) {
      data.category_id = null;
    } else {
      data.category_id = parseInt(data.category_id);
    }
  } else {
    hasError = true
    valid.key.category = true
    valid.message.category = 'カテゴリーを選択してください'
  }
  return [valid, hasError];
}

export default ItemValidation