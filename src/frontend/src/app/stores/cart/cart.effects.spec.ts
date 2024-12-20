import { TestBed } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { MockStore } from '@ngrx/store/testing';
import { Observable, of } from 'rxjs';

import { CartEffects } from './cart.effects';
import { getUser } from '../login/login.selectors';
import * as CartActions from './cart.actions';
import { CartService } from 'src/app/services/cart.service';
import { Cart } from 'src/app/models/cart';
import { stubbedProducts } from 'src/app/unit-test-fixtures/products-utils';
import { completeObservable } from 'src/app/unit-test-fixtures/observable-utils';
import { OrderService } from 'src/app/services/order.service';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';
import { createCart } from 'src/app/unit-test-fixtures/carts-utils';
import { createUser } from 'src/app/unit-test-fixtures/user-utils';

describe('CartEffects', () => {
  let actions$: Observable<any>;
  let effects: CartEffects;
  let myCart: Cart[];
  let getCartSpy: jasmine.Spy<() => Observable<Cart[]>>;
  let addCartSpy: jasmine.Spy<() => Observable<void>>;
  let removeFromCartSpy: jasmine.Spy<() => Observable<void>>;
  let addFromCartSpy: jasmine.Spy<() => Observable<string>>;
  let orderId: string;
  let cartService: CartService;
  let orderService: OrderService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TestSharedModule],
      providers: [
        CartEffects,
        provideMockActions(() => actions$),
      ]
    }).compileComponents();

    myCart = createCart();
    const store = TestBed.inject(MockStore);
    store.overrideSelector(getUser, createUser());
    cartService = TestBed.inject(CartService);
    orderService = TestBed.inject(OrderService);
    getCartSpy = spyOn(cartService, 'getCart').and.returnValue(of(myCart));
    addCartSpy = spyOn(cartService, 'add').and.returnValue(completeObservable());
    removeFromCartSpy = spyOn(cartService, 'delete').and.returnValue(completeObservable());
    orderId = '1';
    addFromCartSpy = spyOn(orderService, 'addFromCart').and.returnValue(of(orderId));
    effects = TestBed.inject(CartEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });

  it('should fetch cart', (done) => {
    // arrange
    const expectedAction = CartActions.fetchCartSuccess({ cart: createCart() });
    
    // act
    actions$ = of(CartActions.fetchCart());

    // assert
    effects.fectchCart$.subscribe(f => {
      expect(f).toEqual(expectedAction);
      done();
    });
  });

  it('fetch cart while the connection is down should perform failed action', (done) => {
    // arrange
    const expectedAction = CartActions.fetchCartFailed({ error: 'Sprawdź połączenie z internetem' });
    getCartSpy.and.returnValue(new Observable(o => o.error({ status: 0 })));
    
    // act
    actions$ = of(CartActions.fetchCart());

    // assert
    effects.fectchCart$.subscribe(f => {
      expect(f).toEqual(expectedAction);
      done();
    });
  });

  it('fetch cart while actions was unsuccessful should perform a failed action', (done) => {
    // arrange
    const expectedAction = CartActions.fetchCartFailed({ error: 'Coś poszło nie tak, spróbuj później' });
    getCartSpy.and.returnValue(new Observable(o => o.error({ status: 500 })));

    // act
    actions$ = of(CartActions.fetchCart());

    // assert
    effects.fectchCart$.subscribe(f => {
      expect(f).toEqual(expectedAction);
      done();
    });

  });

  it('should add product to cart', (done) => {
    // arrange
    const expectedAction = CartActions.addProductToCartSuccess();
    const products = stubbedProducts();

    // act
    actions$ = of(CartActions.addProductToCart({ product: products[0] }));

    effects.addProductToCart$.subscribe(a => {
      expect(a).toEqual(expectedAction);
      done();
    });
  });

  it('add product to cart while connection is down should perform a failed action', (done) => {
    // arrange
    const expectedAction = CartActions.addProductToCartFailed({ error: 'Sprawdź połączenie z internetem' });
    const products = stubbedProducts();
    addCartSpy.and.returnValue(new Observable(o => o.error({ status: 0 })));
    
    // act
    actions$ = of(CartActions.addProductToCart({ product: products[0] }));

    // assert
    effects.addProductToCart$.subscribe(a => {
      expect(a).toEqual(expectedAction);
      done();
    });
  });

  it('add product to cart while action was unsuccessful should perform a failed action', (done) => {
    // arrange
    const expectedAction = CartActions.addProductToCartFailed({ error: 'Coś poszło nie tak, spróbuj później' });
    const products = stubbedProducts();
    addCartSpy.and.returnValue(new Observable(o => o.error({ status: 500 })));
    
    // act
    actions$ = of(CartActions.addProductToCart({ product: products[0] }));

    // assert
    effects.addProductToCart$.subscribe(a => {
      expect(a).toEqual(expectedAction);
      done();
    });
  });

  it('should remove product from cart', (done) => {
    // arrange
    const expectedAction = CartActions.removeProductFromCartSuccess();

    // act
    actions$ = of(CartActions.removeProductFromCart({ cart: myCart[0] }));
    
    // assert
    effects.removeProductFromCart$.subscribe(r => {
      expect(r).toEqual(expectedAction);
      done();
    });
  });

  it('remove from cart while connection is down should perform failed action', (done) => {
    // arrange
    const expectedAction = CartActions.removeProductFromCartFailed({ error: 'Sprawdź połączenie z internetem' });
    removeFromCartSpy.and.returnValue(new Observable(o => o.error({ status: 0 })));

    // act
    actions$ = of(CartActions.removeProductFromCart({ cart: myCart[0] }));
    
    // assert
    effects.removeProductFromCart$.subscribe(r => {
      expect(r).toEqual(expectedAction);
      done();
    });
  });

  it('remove from cart while action was unsuccessful should perform failed action', (done) => {
    // arrange
    const expectedAction = CartActions.removeProductFromCartFailed({ error: 'Coś poszło nie tak, spróbuj później' });
    removeFromCartSpy.and.returnValue(new Observable(o => o.error({ status: 500 })));

    // act
    actions$ = of(CartActions.removeProductFromCart({ cart: myCart[0] }));
    
    // assert
    effects.removeProductFromCart$.subscribe(r => {
      expect(r).toEqual(expectedAction);
      done();
    });
  });

  it('should finalize cart', (done) => {
    // arrange
    const expectedAction = CartActions.finalizeCartSuccess({ orderId });

    // act
    actions$ = of(CartActions.finalizeCart());

    // assert
    effects.finalizeCart$.subscribe(fc => {
      expect(fc).toEqual(expectedAction);
      done();
    });
  });

  it('finalize cart while connection is down should perform failed action', (done) => {
    // arrange
    const expectedAction = CartActions.finalizeCartFailed({ error: 'Sprawdź połączenie z internetem' });
    addFromCartSpy.and.returnValue(new Observable(o => o.error({ status: 0 })));

    // act
    actions$ = of(CartActions.finalizeCart());

    // assert
    effects.finalizeCart$.subscribe(fc => {
      expect(fc).toEqual(expectedAction);
      done();
    });
  });

  it('finalize cart while action was unsuccessful should perform failed action', (done) => {
    // arrange
    const expectedAction = CartActions.finalizeCartFailed({ error: 'Coś poszło nie tak, spróbuj później' });
    addFromCartSpy.and.returnValue(new Observable(o => o.error({ status: 500 })));

    // act
    actions$ = of(CartActions.finalizeCart());

    // assert
    effects.finalizeCart$.subscribe(fc => {
      expect(fc).toEqual(expectedAction);
      done();
    });
  });
});
