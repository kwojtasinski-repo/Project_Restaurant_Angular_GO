import { TestBed } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { provideMockStore, MockStore } from '@ngrx/store/testing';
import { Observable, of } from 'rxjs';

import { CartEffects } from './cart.effects';
import { initialState } from './cart.reducers';
import { initialState as initialLoginState } from '../login/login.reducers';
import { getUser } from '../login/login.selectors';
import { HttpClientModule } from '@angular/common/http';
import { User } from 'src/app/models/user';
import * as CartActions from './cart.actions';
import { CartService } from 'src/app/services/cart.service';
import { Cart } from 'src/app/models/cart';
import { stubbedProducts } from 'src/app/unit-test-fixtures/test-utils';
import { OrderService } from 'src/app/services/order.service';

describe('CartEffects', () => {
  let actions$: Observable<any>;
  let effects: CartEffects;
  let myCart: Cart[];
  let getCartSpy: jasmine.Spy<() => Observable<Cart[]>>;
  let addCartSpy: jasmine.Spy<() => Observable<void>>;
  let removeFromCartSpy: jasmine.Spy<() => Observable<void>>;
  let addFromCartSpy: jasmine.Spy<() => Observable<number>>;
  let orderId: number;
  let cartService: CartService;
  let orderService: OrderService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        CartEffects,
        provideMockActions(() => actions$),
        provideMockStore({ initialState }),
        provideMockStore({ initialState: initialLoginState }),
        {
          provide: "API_URL", useValue: ''
        }
      ],
      imports: [
        HttpClientModule,
      ]
    });

    myCart = createCart();
    const store = TestBed.inject(MockStore);
    store.overrideSelector(getUser, {
      id: 1,
      email: 'email@email',
      role: 'test',
      deleted: null
    } as User);
    cartService = TestBed.inject(CartService);
    orderService = TestBed.inject(OrderService);
    getCartSpy = spyOn(cartService, 'getCart').and.returnValue(of(myCart));
    addCartSpy = spyOn(cartService, 'add').and.returnValue(of());
    removeFromCartSpy = spyOn(cartService, 'delete').and.returnValue(of());
    orderId = 1;
    addFromCartSpy = spyOn(orderService, 'addFromCart').and.returnValue(of(orderId));
    effects = TestBed.inject(CartEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });

  it('should fetch cart', () => {
    // arrange
    const expectedAction = CartActions.fetchCartSuccess({ cart: createCart() });
    
    // act
    actions$ = of(CartActions.fetchCart());

    // assert
    effects.fectchCart$.subscribe(f => expect(f).toEqual(expectedAction));
  });

  it('fetch cart while the connection is down should perform failed action', () => {
    // arrange
    const expectedAction = CartActions.fetchCartFailed({ error: 'Sprawdź połączenie z internetem' });
    getCartSpy.and.returnValue(new Observable(o => o.error({ status: 0 })));
    
    // act
    actions$ = of(CartActions.fetchCart());

    // assert
    effects.fectchCart$.subscribe(f => expect(f).toEqual(expectedAction));

  });

  it('fetch cart while actions was unsuccessful should perform a failed action', () => {
    // arrange
    const expectedAction = CartActions.fetchCartFailed({ error: 'Coś poszło nie tak, spróbuj później' });
    getCartSpy.and.returnValue(new Observable(o => o.error({ status: 500 })));

    // act
    actions$ = of(CartActions.fetchCart());

    // assert
    effects.fectchCart$.subscribe(f => expect(f).toEqual(expectedAction));

  });

  it('should add product to cart', () => {
    // arrange
    const expectedAction = CartActions.addProductToCartSuccess();
    const products = stubbedProducts();

    // act
    actions$ = of(CartActions.addProductToCart({ product: products[0] }));

    effects.addProductToCart$.subscribe(a => expect(a).toEqual(expectedAction))
  });

  it('add product to cart while connection is down should perform a failed action', () => {
    // arrange
    const expectedAction = CartActions.addProductToCartFailed({ error: 'Sprawdź połączenie z internetem' });
    const products = stubbedProducts();
    addCartSpy.and.returnValue(new Observable(o => o.error({ status: 0 })));
    
    // act
    actions$ = of(CartActions.addProductToCart({ product: products[0] }));

    // assert
    effects.addProductToCart$.subscribe(a => expect(a).toEqual(expectedAction));
  });

  it('add product to cart while action was unsuccessful should perform a failed action', () => {
    // arrange
    const expectedAction = CartActions.addProductToCartFailed({ error: 'Coś poszło nie tak, spróbuj później' });
    const products = stubbedProducts();
    addCartSpy.and.returnValue(new Observable(o => o.error({ status: 500 })));
    
    // act
    actions$ = of(CartActions.addProductToCart({ product: products[0] }));

    // assert
    effects.addProductToCart$.subscribe(a => expect(a).toEqual(expectedAction));
  });

  it('should remove product from cart', () => {
    // arrange
    const expectedAction = CartActions.removeProductFromCartSuccess();

    // act
    actions$ = of(CartActions.removeProductFromCart({ cart: myCart[0] }));
    
    // assert
    effects.removeProductFromCart$.subscribe(r => expect(r).toEqual(expectedAction));
  });

  it('remove from cart while connection is down should perform failed action', () => {
    // arrange
    const expectedAction = CartActions.removeProductFromCartFailed({ error: 'Sprawdź połączenie z internetem' });
    removeFromCartSpy.and.returnValue(new Observable(o => o.error({ status: 0 })));

    // act
    actions$ = of(CartActions.removeProductFromCart({ cart: myCart[0] }));
    
    // assert
    effects.removeProductFromCart$.subscribe(r => expect(r).toEqual(expectedAction));
  });

  it('remove from cart while action was unsuccessful should perform failed action', () => {
    // arrange
    const expectedAction = CartActions.removeProductFromCartFailed({ error: 'Coś poszło nie tak, spróbuj później' });
    removeFromCartSpy.and.returnValue(new Observable(o => o.error({ status: 500 })));

    // act
    actions$ = of(CartActions.removeProductFromCart({ cart: myCart[0] }));
    
    // assert
    effects.removeProductFromCart$.subscribe(r => expect(r).toEqual(expectedAction));
  });

  it('should finalize cart', () => {
    // arrange
    const expectedAction = CartActions.finalizeCartSuccess({ orderId });

    // act
    actions$ = of(CartActions.finalizeCart());

    // assert
    effects.finalizeCart$.subscribe(fc => expect(fc).toEqual(expectedAction));
  });

  it('finalize cart while connection is down should perform failed action', () => {
    // arrange
    const expectedAction = CartActions.finalizeCartFailed({ error: 'Sprawdź połączenie z internetem' });
    addFromCartSpy.and.returnValue(new Observable(o => o.error({ status: 0 })));

    // act
    actions$ = of(CartActions.finalizeCart());

    // assert
    effects.finalizeCart$.subscribe(fc => expect(fc).toEqual(expectedAction));
  });

  it('finalize cart while action was unsuccessful should perform failed action', () => {
    // arrange
    const expectedAction = CartActions.finalizeCartFailed({ error: 'Coś poszło nie tak, spróbuj później' });
    addFromCartSpy.and.returnValue(new Observable(o => o.error({ status: 500 })));

    // act
    actions$ = of(CartActions.finalizeCart());

    // assert
    effects.finalizeCart$.subscribe(fc => expect(fc).toEqual(expectedAction));
  });
});

const createCart = () => {
  return [
    {
      id: 1,
      product: {
        id: 1,
        name: 'Product#1',
        description: 'Desc',
        price: 100,
        category: {
          id: 1,
          name: 'Category',
          deleted: false
        },
        deleted: false,
      },
      userId: 1
    }, 
    {
      id: 2,
      product: {
        id: 1,
        name: 'Product#1',
        description: 'Desc',
        price: 100,
        category: {
          id: 1,
          name: 'Category',
          deleted: false
        },
        deleted: false
      },
      userId: 1
    },
    {
      id: 3,
      product: {
        id: 1,
        name: 'Product#1',
        description: 'Desc',
        price: 100,
        category: {
          id: 1,
          name: 'Category',
          deleted: false
        },
        deleted: false
      },
      userId: 1
    }
  ] as Cart[];
}
