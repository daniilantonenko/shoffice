<?php

ini_set('display_errors', '1');
ini_set('display_startup_errors', '1');
error_reporting(E_ALL);

$_POST = json_decode(file_get_contents("php://input"), true);

use PHPMailer\PHPMailer\PHPMailer;
use PHPMailer\PHPMailer\SMTP;
use PHPMailer\PHPMailer\Exception;

require '../php/phpmailer/src/Exception.php';
require '../php/phpmailer/src/PHPMailer.php';
require '../php/phpmailer/src/SMTP.php';

echo 'START ';

if (empty($_POST)) {
    //throw new Exception('Данные отстствуют!');
    exit('Данные отстствуют!');
} else {
    echo 'RUN RUN RUN';
    echo htmlspecialchars($_POST);

    //If everything above is ok then move further 
    $name = $_POST['name'];
    $phone = $_POST['phone'];
    $email = $_POST['email'];
    $message = $_POST['message'];

    //Use PHPmailer
    $mail = new PHPMailer(true);

    //Mailbody to send in email
    $mailbody = '<p><b>Customer contact details:-</b></p>
<p><b>Name:</b> ' . $name . '</p>
<p><b>Email:</b> ' . $email . '</p>
<p><b>Message:</b> ' . $message . '</p>';
}



try {
    //Sender email address
    $sender = SMTP_RECEIVER; // Gmail SMTP username
    $mail->isSMTP(); // Set mailer to use SMTP
    $mail->SMTPDebug = 4; //See errors, change it to 4
    $mail->Host = SMTP_HOST; // Specify main and backup SMTP servers
    $mail->SMTPAuth = true; // Enable SMTP authentication
    $mail->Username = $sender; // SMTP username
    $mail->Password = SMTP_PASSWORD; // Gmail SMTP password
    $mail->SMTPSecure = 'tls'; // Enable TLS encryption, `SSL` also accepted
    $mail->Port = 587; // TCP port to connect to	 

    $mail->setFrom($email, $name);

    $mail->addAddress('daniilantonenko@ya.ru', 'recipientname'); // Add a recipient
//$mail->addAddress('name@example.com'); //add more recipient (optional)
    $mail->addReplyTo($email, $name);
    //$mail->addCC('cc@example.com');
//$mail->addBCC('bcc@example.com');

    //Add attachment 
    if (is_array($_FILES)) {
        $mail->AddAttachment($uploadedFile, $fileName, 'base64', 'mime/type');
    }

    $mail->isHTML(true); // Set email format to HTML
    $mail->Subject = 'Contact Form Submitted by ' . $name;
    $mail->Body = $mailbody;
    //$mail->AltBody = 'This is the body in plain text for non-HTML mail clients';

    //Email your contact form data with an attachment using Gmail SMTP with PHPMailer
    if (!$mail->send()) {
        /*?>
        <div class="alert alert-danger alert-dismissible fade show">
        <strong>Message could not be sent. Please try again.</strong>
        <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
        </div>
        <?php*/
        echo 'Mailer Error: ' . $mail->ErrorInfo;
    } else {
        /*?>
        <div class="alert alert-success alert-dismissible fade show">
        <strong>Message has been sent successfully.</strong>
        <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
        </div>
        <?php*/
        @unlink($uploadedFile);
    }
} catch (Exception $e) {
    echo "Message could not be sent. Mailer Error: {$mail->ErrorInfo}";
}

?>