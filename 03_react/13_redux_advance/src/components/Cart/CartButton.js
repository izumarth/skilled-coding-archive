import { useDispatch, useSelector } from 'react-redux';
import { uiActions } from '../../store/ui-slice';
import classes from './CartButton.module.css';

const CartButton = (props) => {
  const dispath = useDispatch();
  const cartQuantity = useSelector( state => state.cart.totalQuantity);

  const toggleCartHanndler = () => {
    dispath(uiActions.toggle());
  };

  return (
    <button className={classes.button} onClick={toggleCartHanndler}>
      <span>My Cart</span>
      <span className={classes.badge}>{cartQuantity}</span>
    </button>
  );
};

export default CartButton;
