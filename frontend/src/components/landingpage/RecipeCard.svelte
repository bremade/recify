<script>
  import Card, {Content} from '@smui/card';
  import { Row, Col } from 'sveltestrap';
  import Rating from './Stars.svelte';
  import Clock from "svelte-material-icons/Timer.svelte";
  import Euro from "svelte-material-icons/CurrencyEur.svelte";
  import formatter from '../../utils/formatter';
  
  let imageSrc = '/testPic.jpg';

  let rating = Math.floor(Math.random()*(5 - 0) + 0);

  let size = "22px";
  let color = "#2196f3";

  export let recipeId;
  export let recipePrice = 123;
  export let recipeTime = 150;
</script>
<a href="/recipe/{recipeId}">
  <Card>
      <img class="cardImage" src={imageSrc} alt={recipeId}/>
      <Content class="cardContent">
        <h3 style="margin: 0 0 10px;">{recipeId}</h3>
        <Row class="mt-4">
          <Col align="left">
            <div class="flexParent">
              <div class="flexChild">
                <Clock {color} {size}/>
              </div>
              <p>{formatter.formatTime(recipeTime)}</p>
            </div>
          </Col>
          <Col align="center">
            <div class="flexParent">
              <Euro {color} {size}/>
              <p>{formatter.formatPrice(recipePrice)}</p>
            </div>
          </Col>
          <Col align="right">
            <Rating {rating} {size}/>
          </Col>
        </Row>
      </Content>
  </Card>
</a>
