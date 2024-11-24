import { RequestState } from 'src/app/models/request-state';
import * as CartActions from './cart.actions';
import { initialState, cartReducer } from './cart.reducers';
import { CartState } from './cart.state';
import { createCart } from 'src/app/unit-test-fixtures/carts-utils';

describe('CartReducers', () => {
  it('fetch cart', () => {
    // arrange
    const action = CartActions.fetchCart();
    const expectedState: CartState = {
      ...initialState,
      fetchState: RequestState.loading,
      error: null
    };
 
    // act
    const result = cartReducer(initialState, action);
    
    // assert
    expect(result).toEqual(expectedState);
  });

  it('fetch cart success', () => {
    // arrange
    const cart = createCart();
    const action = CartActions.fetchCartSuccess({ cart });
    const expectedState: CartState = {
      ...initialState,
      cart,
      fetchState: RequestState.success,
      error: null
    };
 
    // act
    const result = cartReducer(initialState, action);
    
    // assert
    expect(result).toEqual(expectedState);
  });
});
