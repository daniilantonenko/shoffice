<!doctype html>
<html lang="ru">

<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="description" content="">
  <meta name="author" content="Daniil Antonenko">
  <title>Поделись файлом свом - и он еще не раз вернется.</title>
  <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css" rel="stylesheet"
    integrity="sha384-GLhlTQ8iRABdZLl6O3oVMWSktQOp6b7In1Zl3/Jr59b6EGGoI1aFkw7cmDA6j6gD" crossorigin="anonymous">

  <!--<link href="starter-template.css" rel="stylesheet">-->

</head>

<body>

  <div class="col-lg-8 mx-auto p-4 py-md-5">
    <header class="d-flex align-items-center pb-3 mb-5 border-bottom">
      <a href="/" class="d-flex align-items-center text-dark text-decoration-none">


        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="me-2"
          viewBox="0 0 16 16">
          <title>Поделись</title>
          <path fill-rule="evenodd" clip-rule="evenodd"
            d="M11 2.5a2.5 2.5 0 1 1 .603 1.628l-6.718 3.12a2.499 2.499 0 0 1 0 1.504l6.718 3.12a2.5 2.5 0 1 1-.488.876l-6.718-3.12a2.5 2.5 0 1 1 0-3.256l6.718-3.12A2.5 2.5 0 0 1 11 2.5z"
            fill="currentColor"></path>
        </svg>

        <span class="fs-4">Поделись файлом с офисом</span>
      </a>
    </header>

    <main>
      <!--
    <h1>Get started with Bootstrap</h1>
    <p class="fs-5 col-md-8">Quickly and easily get started with Bootstrap's compiled, production-ready files with this barebones example featuring some basic HTML and helpful links. Download all our examples to get started.</p>

    <div class="mb-5">
      <a href="../examples/" class="btn btn-primary btn-lg px-4">Загрузить</a>
    </div>
    -->

      <form id="sendForm" class="needs-validation" novalidate enctype="multipart/form-data">

        <!--
        <div>
          <div class="mb-3">
            <div class="form-group">
              <input type="text" class="form-control" name="name" placeholder="Имя" required>
            </div>
          </div>
          <div class="mb-3">
            <div class="form-group">
              <input type="email" class="form-control" name="email" placeholder="E-mail" required>
            </div>
          </div>
          <div class="mb-3">
            <div class="form-group">
              <textarea type="text" class="form-control" name="message" placeholder="Сообщение..." rows="5"
                required></textarea>
            </div>
          </div>
        </div>
        -->
        
        <div class="mb-3">
          <div class="form-group">
            <input type="file" class="form-control" name="file" required id="inputFile" multiple
              accept="image/*,.pdf" />
            <div class="invalid-feedback" for="inputFile">
              Необходимо прикрепить один или несколько файлов.
            </div>
          </div>
        </div>

        <div class="mb-3">
          <div class="form-group visually-hidden" id="preview">
            <p>В настоящее время файлы для загрузки не выбраны</p>
          </div>
        </div>

        <div class="mb-3">
          <div class="form-check">
            <input type="checkbox" class="form-check-input" value="" id="152-fz" required id="exampleCheck1">
            <label class="form-check-label" for="invalidCheck">
              Согласен на обработку персональных данных
            </label>
            <div class="invalid-feedback">
              Перед отправкой формы необходимо подтвердить согласие.
            </div>
          </div>
        </div>

        <div class="mb-3">
          <div class="form-group">
            <button type="submit" class="btn btn-primary" id="submit" disabled>Отправить</button>
          </div>
        </div>

      </form>

      <div class="mb-3">
        <div class="form-group" id="endSendForm" style="display: none;">
          <p>Файлы были успешно отправлены!</p>
        </div>
      </div>




      <!--
  <hr class="col-3 col-md-2 mb-5">

    <div class="row g-5">
      <div class="col-md-6">
        <h2>Starter projects</h2>
        <p>Ready to beyond the starter template? Check out these open source projects that you can quickly duplicate to a new GitHub repository.</p>
        <ul class="icon-list ps-0">
          <li class="d-flex align-items-start mb-1"><a href="https://github.com/twbs/bootstrap-npm-starter" rel="noopener" target="_blank">Bootstrap npm starter</a></li>
          <li class="text-muted d-flex align-items-start mb-1">Bootstrap Parcel starter (coming soon!)</li>
        </ul>
      </div>

      <div class="col-md-6">
        <h2>Guides</h2>
        <p>Read more detailed instructions and documentation on using or contributing to Bootstrap.</p>
        <ul class="icon-list ps-0">
          <li class="d-flex align-items-start mb-1"><a href="../getting-started/introduction/">Bootstrap quick start guide</a></li>
          <li class="d-flex align-items-start mb-1"><a href="../getting-started/webpack/">Bootstrap Webpack guide</a></li>
          <li class="d-flex align-items-start mb-1"><a href="../getting-started/parcel/">Bootstrap Parcel guide</a></li>
          <li class="d-flex align-items-start mb-1"><a href="../getting-started/vite/">Bootstrap Vite guide</a></li>
          <li class="d-flex align-items-start mb-1"><a href="../getting-started/contribute/">Contributing to Bootstrap</a></li>
        </ul>
      </div>
    </div>
  </main>
  <footer class="pt-5 my-5 text-muted border-top">
    Created by the Bootstrap team &middot; &copy; 2022
  </footer>
</div>
-->
      <script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/3.6.4/jquery.min.js"></script>
      <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js"
        integrity="sha384-w76AqPfDkMBDXo30jS1Sgez6pr3x5MlQ1ZAGC+nuZB+EYdgRZgiwxhTBTkF7CXvN"
        crossorigin="anonymous"></script>

      <script type="text/javascript" src="scripts.js"></script>



</body>

</html>