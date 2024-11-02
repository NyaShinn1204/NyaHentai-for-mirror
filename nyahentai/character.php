<?php
if (isset($_GET['name'])) {
    $name = $_GET['name'];
}
?>

<?php
function generateRandomString($length = 10)
{
    $characters = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
    $charactersLength = strlen($characters);
    $randomString = '';

    for ($i = 0; $i < $length; $i++) {
        $randomString .= $characters[random_int(0, $charactersLength - 1)];
    }

    return $randomString;
}
?>

<!DOCTYPE html>
<html lang="ja">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <meta name="msapplication-TitleColor" content="#E4E9F7">
    <meta name="theme-color" content="#E4E9F7">
    <meta name="title" content="幻想の森 - お知らせ">
    <meta name="description" content="幻想の森で数ある有名の作品を読もう！">
    <meta property="og:type" content="website">
    <meta property="og:url" content="https://dev.minohaed.com">
    <meta property="og:title" content="幻想の森 - お知らせ">
    <meta property="og:description" content="幻想の森で数ある有名の作品を無料で読もう！">
    <meta name="keywords" content="無料で漫画,raw,raw-free,漫画,漫画村,幻想の森">

    <script type="text/javascript" src="/assets/js/darkmode.js"></script>

    <title>幻想の森 - お知らせ</title>

    <link rel="stylesheet" href="/assets/style/font-awesome.css?version=<?php echo generateRandomString(16) ?>">
    <link rel="stylesheet" href="/assets/style/web-style.css?version=<?php echo generateRandomString(16) ?>">
    <link rel="stylesheet" href="/assets/style/web-drawer.css?version=<?php echo generateRandomString(16) ?>">
    <link rel="stylesheet" href="/assets/style/web-font.css?version=<?php echo generateRandomString(16) ?>">
    <link rel="stylesheet" href="/assets/style/web-header.css?version=<?php echo generateRandomString(16) ?>">
    <link rel="stylesheet" href="/assets/style/web-doujinshi.css?version=<?php echo generateRandomString(16) ?>">
    <link rel="stylesheet" href="/assets/style/doujinshi/slick.css?version=<?php echo generateRandomString(16) ?>">
    <link rel="stylesheet" href="/assets/style/doujinshi/slick-theme.css?version=<?php echo generateRandomString(16) ?>">

    <script type="text/javascript" src="https://code.jquery.com/jquery-1.11.0.min.js"></script>
    <script type="text/javascript" src="https://code.jquery.com/jquery-migrate-1.2.1.min.js"></script>
    <script type="text/javascript" src="https://cdn.jsdelivr.net/npm/slick-carousel@1.8.1/slick/slick.min.js"></script>
    <script type="text/javascript" src="/assets/js/global.js?version=<?php echo generateRandomString(16) ?>"></script>
</head>

<body>
    <div class="header">
        <button class="drawer-toggle">☰</button>
        <a href="/">
            <h1>幻想の森</h1>
        </a>
        <div class="left-header">
            <div id="mobile_search"><i class="fa fa-search"></i></div>
            <div id="search">
                <div class="search-content">
                    <form action="/search" search_form="true" method="get" autocomplete="off">
                        <input type="text" name="text_param" class="form-control search-input" placeholder="キーワードを入力してください...">
                        <button type="submit" class="search-icon"><i class="fa fa-search"></i></button>
                    </form>
                    <div class="nav search-result-pop" id="search-suggest">
                        <div class="loading-relative" id="search-loading" style="min-height:60px;display: none;">
                            <div class="loading">
                                <div class="span1"></div>
                                <div class="span2"></div>
                                <div class="span3"></div>
                            </div>
                        </div>
                        <div class="result" style="display:none;"></div>
                    </div>
                </div>
            </div>
            <button class="mobile_dark-mode-toggle" id="dark-mode-toggle">🌓</button>
            <button class="dark-mode-toggle" id="dark-mode-toggle">🌓ダークモード</button>
        </div>
    </div>
    <div class="drawer-menu">
        <ul>
            <li>
                <button class="close-btn">
                    <i class="fa fa-angle-left" aria-hidden="true"></i>
                    <!--<span class="close-btn-text">閉じる</span-->
                </button>
                <span class="close-btn-text">閉じる</span>
            </li>
            <li><a href="/">ホーム</a></li>
            <li><a href="#">お知らせ</a></li>
            <li><a href="/requests">リクエスト</a></li>
            <li><a href="/search/">検索</a></li>
            <li><a href="#">なんかつくる</a></li>
            <li><a href="#">なんかつくる</a></li>
            <li><a href="#">なんかつくる</a></li>
        </ul>
    </div>
    <div class="container_doujin">
        <header id="header">
            <div id="header-in">

                <div id="logo"><a href="https://dev.minohaed.com/doujinshi/nyahentai/"><i>-</i><span>Nya<i>:</i></span>Hentai</a><i>-</i></div>

                <input id="menu-input" type="checkbox" class="menu-unshown">
                <label id="menu-open" for="menu-input"></label>
                <label id="menu-close" for="menu-input" class="menu-unshown"></label>
                <div id="menu-content">
                    <label id="menu-close" for="menu-input" class="menu-unshown"></label>
                    <div id="menu-menu">
                        <h2>メニュー</h2>
                        <a href="https://nyahentai.re/mylist/">マイリスト</a>
                        <a href="https://nyahentai.re/fanzine/">同人誌</a>
                        <a href="https://nyahentai.re/magazine/">商業誌</a>
                        <a href="https://nyahentai.re/rising/">急上昇</a>
                        <a href="https://nyahentai.re/popularity/">人気</a>
                        <a href="https://nyahentai.re/contact/">お問い合わせ</a>
                    </div>
                </div>

                <input id="search-input" type="checkbox" class="search-unshown">
                <label id="search-open" for="search-input"></label>
                <label id="search-close" for="search-input" class="search-unshown"></label>
                <div id="search-content">
                    <label id="search-close" for="search-input" class="search-unshown"></label>
                    <div id="search-search">
                        <h2>検索</h2>
                        <form method="get" action="https://nyahentai.re/" class="search-box">
                            <input type="text" name="s">
                            <input type="submit" value="">
                        </form>
                    </div>
                    <div id="search-tag">
                        <h2>タグから探す</h2>
                        <a href="https://nyahentai.re/parody/">パロディ</a>
                        <a href="https://nyahentai.re/character/">キャラクター</a>
                        <a href="https://nyahentai.re/genre/">ジャンル</a>
                        <a href="https://nyahentai.re/circle/">サークル</a>
                        <a href="https://nyahentai.re/artist/">作者</a>
                    </div>
                </div>

            </div>
        </header>
        <div id="container">
            <div id="main">
                <?php
                function fetchCircle($name)
                {
                    $url = 'https://dev.minohaed.com/api/v1/doujinshi/nyahentai/get_character?name=' . $name;

                    // cURLを使ってAPIリクエストを送信
                    $ch = curl_init($url);
                    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
                    curl_setopt($ch, CURLOPT_HTTPHEADER, [
                        'Content-Type: application/json',
                    ]);

                    $response = curl_exec($ch);
                    curl_close($ch);

                    // レスポンスをJSONとしてデコード
                    return json_decode($response, true);
                }

                // APIからデータを取得
                $circleList = fetchCircle($name);
                ?>
                <div class="page-h">
                    <h3>キャラクター</h3>
                    <h1><?php echo $circleList["h1"]; ?></h1>
                    <span><?php echo $circleList["span"]; ?></span>
                </div>
                <?php
                function fetchDoujinshi($name)
                {
                    $url = 'https://dev.minohaed.com/api/v1/doujinshi/nyahentai/get_character_list?name=' . $name;

                    // cURLを使ってAPIリクエストを送信
                    $ch = curl_init($url);
                    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
                    curl_setopt($ch, CURLOPT_HTTPHEADER, [
                        'Content-Type: application/json',
                    ]);

                    $response = curl_exec($ch);
                    curl_close($ch);

                    // レスポンスをJSONとしてデコード
                    return json_decode($response, true);
                }

                // APIからデータを取得
                $doujinshiList = fetchDoujinshi($name);
                ?>
                <div class="post-list">
                    <?php if ($doujinshiList): ?>
                        <?php foreach ($doujinshiList as $item): ?>
                            <a href="<?php echo str_replace("https://nyahentai.re/", "https://dev.minohaed.com/doujinshi/nyahentai/", htmlspecialchars($item['a_href'], ENT_QUOTES, 'UTF-8')); ?>">
                                <div class="post-list-image">
                                    <img src="<?php echo htmlspecialchars($item['img_src'], ENT_QUOTES, 'UTF-8'); ?>" loading="lazy" alt="">
                                </div>
                                <span><?php echo htmlspecialchars($item['span_text'], ENT_QUOTES, 'UTF-8'); ?></span>
                            </a>
                        <?php endforeach; ?>
                    <?php else: ?>
                        <p>データがありません。</p>
                    <?php endif; ?>
                </div>


                <div class="home-h">
                    <h3>parody</h3>
                    <h1>パロディ</h1>
                    <a href="https://nyahentai.re/parody/">more</a>
                </div>
                <?php
                function fetchParody()
                {
                    $url = 'https://dev.minohaed.com/api/v1/doujinshi/nyahentai/tag/get_parody';

                    // cURLを使ってAPIリクエストを送信
                    $ch = curl_init($url);
                    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
                    curl_setopt($ch, CURLOPT_HTTPHEADER, [
                        'Content-Type: application/json',
                    ]);

                    $response = curl_exec($ch);
                    curl_close($ch);

                    // レスポンスをJSONとしてデコード
                    return json_decode($response, true);
                }

                // APIからデータを取得
                $doujinshiList = fetchParody();
                ?>

                <div class="tag-list">
                    <?php if ($doujinshiList): ?>
                        <?php foreach ($doujinshiList as $item): ?>
                            <a href="<?php echo htmlspecialchars($item['a_href'], ENT_QUOTES, 'UTF-8'); ?>">
                                <?php echo htmlspecialchars($item['tag_text'], ENT_QUOTES, 'UTF-8'); ?><span><?php echo htmlspecialchars($item['span_text'], ENT_QUOTES, 'UTF-8'); ?></span>
                            </a>
                        <?php endforeach; ?>
                    <?php else: ?>
                        <p>データがありません。</p>
                    <?php endif; ?>
                </div>
                <div class="home-h">
                    <h3>character</h3>
                    <h1>キャラクター</h1>
                    <a href="https://nyahentai.re/character/">more</a>
                </div>
                <?php
                function fetchCharactor()
                {
                    $url = 'https://dev.minohaed.com/api/v1/doujinshi/nyahentai/tag/get_charactor';

                    // cURLを使ってAPIリクエストを送信
                    $ch = curl_init($url);
                    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
                    curl_setopt($ch, CURLOPT_HTTPHEADER, [
                        'Content-Type: application/json',
                    ]);

                    $response = curl_exec($ch);
                    curl_close($ch);

                    // レスポンスをJSONとしてデコード
                    return json_decode($response, true);
                }

                // APIからデータを取得
                $doujinshiList = fetchCharactor();
                ?>

                <div class="tag-list">
                    <?php if ($doujinshiList): ?>
                        <?php foreach ($doujinshiList as $item): ?>
                            <a href="<?php echo htmlspecialchars($item['a_href'], ENT_QUOTES, 'UTF-8'); ?>">
                                <?php echo htmlspecialchars($item['tag_text'], ENT_QUOTES, 'UTF-8'); ?><span><?php echo htmlspecialchars($item['span_text'], ENT_QUOTES, 'UTF-8'); ?></span>
                            </a>
                        <?php endforeach; ?>
                    <?php else: ?>
                        <p>データがありません。</p>
                    <?php endif; ?>
                </div>
            </div>
        </div>
    </div>
    <footer>
        <span>
            &copy;2024 <a href="https://nyahentai.re/">NyaHentai</a>
        </span>
    </footer>
    <script type="text/javascript" src="https://nyahentai.re/wp-content/themes/ReHentai/js/headroom.min.js"></script>
    <script>
        (function() {
            var header = document.querySelector(".container_doujin #header");
            if (window.location.hash) {
                header.classList.add("headroom--unpinned");
            }
            var headroom = new Headroom(header, {
                tolerance: {
                    down: 10,
                    up: 20
                },
                offset: 205
            });
            headroom.init();
        }());
    </script>
    </div>
</body>

</html>