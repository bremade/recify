<script>
  import Card, {Content, PrimaryAction, Media, MediaContent, Actions, ActionButtons, ActionIcons} from '@smui/card';
  import { Container, Row, Col } from 'sveltestrap';
  
  let imageSrc = '/testPic.jpg';

  export let recipeId;
</script>
<a href="/recipe/{recipeId}">
  <Card>
      <img class="cardImage" src={imageSrc} alt={recipeId}/>
      <Content class="cardContent">
        <h3 style="margin: 0 0 10px;">{recipeId}</h3>
        <Row>
          <Col align="left">
            Time
          </Col>
          <Col align="center">
            Price
          </Col>
          <Col align="right">
            Rating
          </Col>
        </Row>
      </Content>
  </Card>
</a>
