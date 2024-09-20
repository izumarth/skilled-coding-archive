import { uiActions } from './ui-slice';
import { cartActions } from './cart-slice';

export const fetchCartData = () => {
  return async (dispatch) => {
    const fetchDate = async () => {
      const response = await fetch('https://hogehoge.firebaseio.com/cart.json');

      if (!response.ok) {
        throw new Error('Fetch cart data failed');
      }

      const data = await response.json();

      return data;
    };

    try {
      const cartData = await fetchDate();
      dispatch(cartActions.replaceCart({
        items: cartData.items || [],
        totalQuantity: cartData.totalQuantity,
      }));
    } catch (error) {
      dispatch(
        uiActions.showNotification({
          status: 'error',
          title: 'Error',
          message: 'Error Fetch cart data success',
        })
      );
    }
  };
};

export const sendCartData = (cart) => {
  return async (dispatch) => {
    dispatch(
      uiActions.showNotification({
        status: 'pending',
        title: 'Sending...',
        message: 'Sending cart Data',
      })
    );

    const sendRequest = async () => {
      const response = await fetch('https://hogehoge.firebaseio.com/cart.json', {
        method: 'PUT',
        body: JSON.stringify({items: cart.items, totalQuantity: cart.totalQuantity}),
      });

      if (!response.ok) {
        throw new Error('Sending cart data failed');
      }
    };

    try {
      await sendRequest();

      dispatch(
        uiActions.showNotification({
          status: 'succcess',
          title: 'Success!!',
          message: 'Send cart data success',
        })
      );
    } catch (error) {
      dispatch(
        uiActions.showNotification({
          status: 'error',
          title: 'Error',
          message: 'Error Send cart data success',
        })
      );
    }
  };
};
